package router

import (
	"fmt"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/handlers"
	"github.com/ghmeier/bloodlines/workers"
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
	sendgrid := gateways.NewSendgrid(config.Sendgrid)
	towncenter := gateways.NewTownCenter(config.TownCenter)

	rabbit, err := gateways.NewRabbit(config.Rabbit)
	if err != nil {
		fmt.Println("ERROR: coud not connect to RabbitMQ")
		fmt.Println(err.Error())
		return nil, err
	}

	b := &Bloodlines{
		content:    handlers.NewContent(sql),
		receipt:    handlers.NewReceipt(sql, sendgrid, towncenter, rabbit),
		job:        handlers.NewJob(sql),
		trigger:    handlers.NewTrigger(sql, sendgrid, towncenter, rabbit),
		preference: handlers.NewPreference(sql),
		workers:    []workers.Send{workers.NewSend(sql, sendgrid, towncenter, rabbit)},
	}

	InitRouter(b)
	return b, nil
}

/*InitRouter connects the handlers to endpoints with gin*/
func InitRouter(b *Bloodlines) {
	b.router = gin.Default()

	content := b.router.Group("/api/content")
	{
		content.POST("", b.content.New)
		content.GET("", b.content.ViewAll)
		content.GET("/:contentId", b.content.View)
		content.PUT("/:contentId", b.content.Update)
		content.DELETE("/:contentId", b.content.Deactivate)
	}
	receipt := b.router.Group("/api/receipt")
	{
		receipt.GET("", b.receipt.ViewAll)
		receipt.POST("/send", b.receipt.Send)
		receipt.GET("/:receiptId", b.receipt.View)
	}
	job := b.router.Group("/api/job")
	{
		job.GET("", b.job.ViewAll)
		job.POST("", b.job.New)
		job.GET("/:jobId", b.job.View)
		job.PUT("/:jobId", b.job.Update)
		job.DELETE("/:jobId", b.job.Stop)
	}
	trigger := b.router.Group("/api/trigger")
	{
		trigger.POST("", b.trigger.New)
		trigger.GET("", b.trigger.ViewAll)
		trigger.GET("/:key", b.trigger.View)
		trigger.PUT("/:key", b.trigger.Update)
		trigger.DELETE("/:key", b.trigger.Remove)
		trigger.POST("/:key/activate", b.trigger.Activate)
	}
	pref := b.router.Group("/api/preference")
	{
		pref.POST("", b.preference.New)
		pref.GET("/:userId", b.preference.View)
		pref.PATCH("/:userId", b.preference.Update)
		pref.DELETE("/:userId", b.preference.Deactivate)
	}

	for _, w := range b.workers {
		w.Consume()
	}
}

/*Start begins the Bloodlines server*/
func (b *Bloodlines) Start(port string) {
	b.router.Run(port)
}
