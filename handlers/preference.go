package handlers

import (
	"github.com/pborman/uuid"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/gateways"
	"github.com/ghmeier/bloodlines/helpers"
	"github.com/ghmeier/bloodlines/models"
)

/*PreferenceI describes the handler for preferences routes*/
type PreferenceI interface {
	New(ctx *gin.Context)
	View(ctx *gin.Context)
	Update(ctx *gin.Context)
	Deactivate(ctx *gin.Context)
}

/*Preference implements PreferenceI for handling requests*/
type Preference struct {
	Helper helpers.PreferenceI
}

/*NewPreference constructs and returns a new preference handler*/
func NewPreference(sql gateways.SQL) *Preference {
	return &Preference{Helper: helpers.NewPreference(sql)}
}

/*New creates and stores a new preference entry based on the given preference struct*/
func (p *Preference) New(ctx *gin.Context) {
	var json models.Preference
	err := ctx.BindJSON(&json)
	if err != nil {
		UserError(ctx, "Error: Unable to parse json", err)
		return
	}

	preference := models.NewPreference(json.UserID)
	err = p.Helper.Insert(preference)
	if err != nil {
		ServerError(ctx, err, json)
		return
	}

	Success(ctx, preference)
}

/*View returns a preference entity associated with the given user id*/
func (p *Preference) View(ctx *gin.Context) {
	id := ctx.Param("userId")

	preference, err := p.Helper.GetByUserID(id)
	if err != nil {
		ServerError(ctx, err, nil)
		return
	}

	Success(ctx, preference)
}

/*Update overwrites a preference entity associated with the given user id*/
func (p *Preference) Update(ctx *gin.Context) {
	id := ctx.Param("userId")

	var json models.Preference
	err := ctx.BindJSON(&json)
	if err != nil {
		UserError(ctx, "Error: unable to parse json", err)
		return
	}

	json.UserID = uuid.Parse(id)
	err = p.Helper.Update(&json)
	if err != nil {
		ServerError(ctx, err, json)
		return
	}

	Success(ctx, json)
}

/*Deactivate sets a user's preference entity to UNSUBSCRIBED*/
func (p *Preference) Deactivate(ctx *gin.Context) {
	id := ctx.Param("userId")

	preference, err := p.Helper.GetByUserID(id)
	if err != nil {
		ServerError(ctx, err, nil)
		return
	}
	preference.Email = models.UNSUBSCRIBED

	err = p.Helper.Update(preference)
	if err != nil {
		ServerError(ctx, err, preference)
		return
	}

	Success(ctx, preference)
}
