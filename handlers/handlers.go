package handlers

import (
	"os"
	"strconv"

	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-contrib/cors.v1"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/dgrijalva/jwt-go"

	"github.com/ghmeier/bloodlines/gateways"
	coi "github.com/ghmeier/coinage/gateways"
	t "github.com/jakelong95/TownCenter/gateways"
	w "github.com/lcollin/warehouse/gateways"
	cov "github.com/yuderekyu/covenant/gateways"
)

/*BaseHandler contains wrapper methods that all handlers need and should use
  for consistency across services*/
type BaseHandler struct {
	Stats *statsd.Client
}

type ExpressoClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

/*GatewayContext contains references to each type of gateway used for simple
  use in handler construction*/
type GatewayContext struct {
	Sql        gateways.SQL
	Sendgrid   gateways.SendgridI
	TownCenter t.TownCenterI
	Covenant   cov.Covenant
	Warehouse  w.Warehouse
	Bloodlines gateways.Bloodlines
	Coinage    coi.Coinage
	Rabbit     gateways.RabbitI
	Stats      *statsd.Client
	Stripe     coi.Stripe
	S3         gateways.S3
	Shippo     w.Shippo
}

/*NewBaseHandler returns a new BaseHandler instance from a given stats*/
func NewBaseHandler(stats *statsd.Client) *BaseHandler {
	return &BaseHandler{Stats: stats}
}

/*GetPaging returns the offset and limit parameters from a gin request context
defaults to offset=0 and limit=20*/
func (b *BaseHandler) GetPaging(ctx *gin.Context) (int, int) {
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	return offset, limit
}

/*UserError sends a 400 response with the given message string and error object*/
func (b *BaseHandler) UserError(ctx *gin.Context, msg string, obj interface{}) {
	if b.Stats != nil {
		b.Stats.Increment("400")
	}
	b.send(ctx, 400, &gin.H{"success": false, "message": msg, "data": obj})
}

/*NotFoundError sends a 404 response and false success when a resource is not present*/
func (b *BaseHandler) NotFoundError(ctx *gin.Context, msg string) {
	if b.Stats != nil {
		b.Stats.Increment("404")
	}
	b.send(ctx, 404, &gin.H{"success": false, "message": msg})
}

/*Unauthorized sends a 401 response along with a message*/
func (b *BaseHandler) Unauthorized(ctx *gin.Context, msg string) {
	if b.Stats != nil {
		b.Stats.Increment("401")
	}
	b.send(ctx, 401, &gin.H{"success": false, "message": msg})
}

/*ServerError sends a 500 response with the given error and object*/
func (b *BaseHandler) ServerError(ctx *gin.Context, err error, obj interface{}) {
	if b.Stats != nil {
		b.Stats.Increment("500")
	}
	b.send(ctx, 500, &gin.H{"success": false, "message": err.Error(), "data": obj})
}

/*Success sends a 200 response with the given object*/
func (b *BaseHandler) Success(ctx *gin.Context, obj interface{}) {
	if b.Stats != nil {
		b.Stats.Increment("200")
	}
	b.send(ctx, 200, &gin.H{"success": true, "data": obj})
}

func (b *BaseHandler) send(ctx *gin.Context, status int, json *gin.H) {
	ctx.JSON(status, json)
}

/*Time sets up gin middleware for sending timing stats*/
func (b *BaseHandler) Time() gin.HandlerFunc {
	return func(c *gin.Context) {
		if b.Stats != nil {
			defer b.Stats.NewTiming().Send(c.Request.Method)
		}
		c.Next()
	}
}

/*GetCors returns a gin handlerFunc for CORS reuquests in expresso services */
func GetCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AddAllowMethods("DELETE")
	config.AddAllowMethods("PATCH")
	config.AddAllowHeaders("X-Auth")
	config.AddExposeHeaders("X-Auth")
	config.AddAllowHeaders("X-Token")
	config.AddExposeHeaders("X-Token")
	config.AllowAllOrigins = true
	return cors.New(config)
}

/*GetJWT returns a gin handlerfunc for authenticating JWTs in expresso services*/
func (b *BaseHandler) GetJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if gin.Mode() == gin.TestMode || gin.Mode() == gin.DebugMode {
			ctx.Next()
			return
		}

		jwtToken := os.Getenv("JWT_TOKEN")
		tokenHeader := ctx.Request.Header.Get("X-Token")
		if tokenHeader != "" && tokenHeader == jwtToken {
			ctx.Next()
			return
		}

		authHeader := ctx.Request.Header.Get("X-Auth")
		token, err := jwt.ParseWithClaims(authHeader, &ExpressoClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtToken), nil
		})

		if err != nil {
			b.Unauthorized(ctx, "Unable to parse token")
			ctx.Abort()
			return
		}

		claims, _ := token.Claims.(*ExpressoClaims)
		if err != nil || !token.Valid || claims.Valid() != nil {
			b.Unauthorized(ctx, "Invalid token")
			ctx.Abort()
			return
		}

		ctx.Request.Header.Add("X-UserId", claims.UserID)
		ctx.Next()
	}
}
