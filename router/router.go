package router

import (
	"fmt"

	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/handlers"
	"github.com/ghmeier/bloodlines/workers"
	tcg "github.com/jakelong95/TownCenter/gateways"
)

/*Bloodlines is the main server object which routes requests*/
type Bloodlines struct {
	router     *gin.Engine
	content    handlers.ContentI
	receipt    handlers.ReceiptI
	job        handlers.JobI
	trigger    handlers.TriggerI
	preference handlers.PreferenceI
	workers    []workers.Send
}

/*New returns a ready-to-run Bloodlines struct from the given config*/
func New(config *config.Root) (*Bloodlines, error) {
	sql, err := gateways.NewSQL(config.SQL)
	if err != nil {
		fmt.Println("ERROR: could not connect to mysql.")
		fmt.Println(err.Error())
		return nil, err
	}

	stats, err := statsd.New(
		statsd.Address(config.Statsd.Host+":"+config.Statsd.Port),
		statsd.Prefix(config.Statsd.Prefix),
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	sendgrid := gateways.NewSendgrid(config.Sendgrid)
	towncenter := tcg.NewTownCenter(config.TownCenter)

	rabbit, err := gateways.NewRabbit(config.Rabbit)
	if err != nil {
		fmt.Println("ERROR: coud not connect to RabbitMQ")
		fmt.Println(err.Error())
	}

	ctx := &handlers.GatewayContext{
		Sql:        sql,
		Sendgrid:   sendgrid,
		TownCenter: towncenter,
		Rabbit:     rabbit,
		Stats:      stats,
	}

	b := &Bloodlines{
		content:    handlers.NewContent(ctx),
		receipt:    handlers.NewReceipt(ctx),
		job:        handlers.NewJob(ctx),
		trigger:    handlers.NewTrigger(ctx),
		preference: handlers.NewPreference(ctx),
		workers:    []workers.Send{workers.NewSend(ctx)},
	}

	InitRouter(b)
	return b, nil
}

/*InitRouter connects the handlers to endpoints with gin*/
func InitRouter(b *Bloodlines) {
	b.router = gin.Default()
	b.router.Use(handlers.GetCors())

	content := b.router.Group("/api/content")
	{
		content.Use(b.content.Time())
		content.POST("", b.content.New)
		content.GET("", b.content.ViewAll)
		content.GET("/:contentId", b.content.View)
		content.PUT("/:contentId", b.content.Update)
		content.DELETE("/:contentId", b.content.Deactivate)
	}

	receipt := b.router.Group("/api/receipt")
	{
		receipt.Use(b.receipt.Time())
		receipt.GET("", b.receipt.ViewAll)
		receipt.POST("/send", b.receipt.Send)
		receipt.GET("/:receiptId", b.receipt.View)
	}

	job := b.router.Group("/api/job")
	{
		job.Use(b.job.Time())
		job.GET("", b.job.ViewAll)
		job.POST("", b.job.New)
		job.GET("/:jobId", b.job.View)
		job.PUT("/:jobId", b.job.Update)
		job.DELETE("/:jobId", b.job.Stop)
	}

	trigger := b.router.Group("/api/trigger")
	{
		trigger.Use(b.trigger.Time())
		trigger.POST("", b.trigger.New)
		trigger.GET("", b.trigger.ViewAll)
		trigger.GET("/:key", b.trigger.View)
		trigger.PUT("/:key", b.trigger.Update)
		trigger.DELETE("/:key", b.trigger.Remove)
		trigger.POST("/:key/activate", b.trigger.Activate)
	}

	pref := b.router.Group("/api/preference")
	{
		pref.Use(b.preference.Time())
		pref.POST("", b.preference.New)
		pref.GET("/:userId", b.preference.View)
		pref.PATCH("/:userId", b.preference.Update)
		pref.DELETE("/:userId", b.preference.Deactivate)
	}

	for _, w := range b.workers {
		w.Consume()
	}
}

func Time(s *statsd.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer s.NewTiming().Send("duration")
		c.Next()
	}
}

/*Start begins the Bloodlines server*/
func (b *Bloodlines) Start(port string) {
	b.router.Run(port)
}
