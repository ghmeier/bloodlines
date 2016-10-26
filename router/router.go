package router

import (
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/handlers"
)

type Bloodlines struct {
	router *gin.Engine
	content handlers.ContentI
	//handlers.Receipt
	//handlers.Job
	//handlers.Trigger
	//handlers.Preference
}

func New() *Bloodlines{
	b := &Bloodlines{
		content: &handlers.Content{},
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

	return b
}

func (b *Bloodlines) Start(port string) {
	b.router.Run(port)
}