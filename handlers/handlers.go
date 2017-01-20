package handlers

import (
	"strconv"

	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/gateways"
)

/*BaseHandler contains wrapper methods that all handlers need and should use
  for consistency across services*/
type BaseHandler struct {
	Stats *statsd.Client
}

/*GatewayContext contains references to each type of gateway used for simple
  use in handler construction*/
type GatewayContext struct {
	Sql        gateways.SQL
	Sendgrid   gateways.SendgridI
	TownCenter gateways.TownCenterI
	Rabbit     gateways.RabbitI
	Stats      *statsd.Client
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
	ctx.JSON(400, &gin.H{"success": false, "message": msg, "error": obj})
}

/*ServerError sends a 500 response with the given error and object*/
func (b *BaseHandler) ServerError(ctx *gin.Context, err error, obj interface{}) {
	if b.Stats != nil {
		b.Stats.Increment("500")
	}
	ctx.JSON(500, &gin.H{"success": false, "message": err.Error(), "error": err, "data": obj})
}

/*Success sends a 200 response with the given object*/
func (b *BaseHandler) Success(ctx *gin.Context, obj interface{}) {
	if b.Stats != nil {
		b.Stats.Increment("200")
	}
	ctx.JSON(200, gin.H{"success": true, "data": obj})
}
