package handlers

import (
	"strconv"

	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/gateways"
)

type BaseHandler struct {
	Stats *statsd.Client
}

type GatewayContext struct {
	Sql        gateways.SQL
	Sendgrid   gateways.SendgridI
	TownCenter gateways.TownCenterI
	Rabbit     gateways.RabbitI
	Stats      *statsd.Client
}

func NewBaseHandler(stats *statsd.Client) *BaseHandler {
	return &BaseHandler{Stats: stats}
}

func (b *BaseHandler) GetPaging(ctx *gin.Context) (int, int) {
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	return offset, limit
}

func (b *BaseHandler) UserError(ctx *gin.Context, msg string, obj interface{}) {
	if b.Stats != nil {
		b.Stats.Increment("400")
	}
	ctx.JSON(400, &gin.H{"success": false, "message": msg, "error": obj})
}

func (b *BaseHandler) ServerError(ctx *gin.Context, err error, obj interface{}) {
	if b.Stats != nil {
		b.Stats.Increment("500")
	}
	ctx.JSON(500, &gin.H{"success": false, "message": err.Error(), "error": obj})
}

func (b *BaseHandler) Success(ctx *gin.Context, obj interface{}) {
	if b.Stats != nil {
		b.Stats.Increment("200")
	}
	ctx.JSON(200, gin.H{"success": true, "data": obj})
}
