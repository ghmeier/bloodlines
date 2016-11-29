package handlers

import (
	"github.com/pborman/uuid"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/helpers"
	"github.com/ghmeier/bloodlines/models"
)

type PreferenceI interface {
	New(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Deactivate(ctx *gin.Context)
}

type Preference struct {
	helper *helpers.Preference
}

func NewPreference(sql gateways.Sql) *Preference {
	return &Preference{helper: helpers.NewPreference(sql)}
}

func (p *Preference) New(ctx *gin.Context) {
	var json models.Preference
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	preference := models.NewPreference(json.UserId)
	err = p.helper.Insert(preference)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": preference})
}

func (p *Preference) View(ctx *gin.Context) {
	id := ctx.Param("userId")
	if id == "" {
		ctx.JSON(400, errResponse("Invalid userId"))
		return
	}

	preference, err := p.helper.GetPreferenceByUserId(id)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": preference})
}

func (p *Preference) Update(ctx *gin.Context) {
	id := ctx.Param("userId")
	if id == "" {
		ctx.JSON(400, errResponse("Invalid userId"))
		return
	}

	var json models.Preference
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	json.UserId = uuid.Parse(id)
	err = p.helper.Update(&json)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}
	ctx.JSON(200, gin.H{"data": json})
}

func (p *Preference) Deactivate(ctx *gin.Context) {
	id := ctx.Param("userId")
	if id == "" {
		ctx.JSON(400, errResponse("Invalid userId"))
		return
	}

	preference := &models.Preference{
		UserId: uuid.Parse(id),
		Email:  models.UNSUBSCRIBED,
	}
	err := p.helper.Update(preference)
	if err != nil {
		ctx.JSON(500, errResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": preference})
}
