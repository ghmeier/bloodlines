package router

import (
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/handlers"
)

type Bloodlines struct {
	router 	   *gin.Engine
	content    handlers.ContentI
	receipt    handlers.ReceiptI
	job 	   handlers.JobI
	trigger    handlers.TriggerI
	preference handlers.PreferenceI
}

func New() *Bloodlines{
	b := &Bloodlines{
		content: 	&handlers.Content{},
		receipt: 	&handlers.Receipt{},
		job:	 	&handlers.Job{},
		trigger:  	&handlers.Trigger{},
		preference: &handlers.Preference{},
	}
	b.router = gin.Default()

	content := b.router.Group("/api/content")
	{
		content.POST("",b.content.New)
		content.GET("",b.content.ViewAll)
		content.GET("/:contentId", b.content.View)
		content.PUT("/:contentId", b.content.Update)
		content.DELETE("/:contentId", b.content.Deactivate)
	}
	receipt := b.router.Group("/api/receipt")
	{
		receipt.GET("",b.receipt.ViewAll)
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
		trigger.GET("/:triggerKey", b.trigger.View)
		trigger.PUT("/:triggerKey", b.trigger.Update)
		trigger.DELETE("/:triggerKey", b.trigger.Remove)
		trigger.POST("/:triggerKey/activate")
	}
	pref := b.router.Group("/api/preference")
	{
		pref.POST("", b.preference.New)
		pref.GET("/:userId", b.preference.View)
		pref.PATCH("/:userId", b.preference.Update)
		pref.DELETE("/:userId", b.preference.Deactivate)
	}

	return b
}

func (b *Bloodlines) Start(port string) {
	b.router.Run(port)
}