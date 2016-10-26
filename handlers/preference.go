package handlers

import (
	"gopkg.in/gin-gonic/gin.v1"
)

type PreferenceI interface {
	New(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Deactivate(ctx *gin.Context)
}

type Preference struct {}

func (p *Preference) New(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (p *Preference) View(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (p *Preference) Update(ctx *gin.Context) {
	ctx.JSON(200, empty())
}

func (p *Preference) Deactivate(ctx *gin.Context) {
	ctx.JSON(200, empty())
}
