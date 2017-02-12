package gateways

import (
	"encoding/json"
	//"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/models"
)

type BloodlinesSuite struct {
	suite.Suite
	bloodlines Bloodlines
	url        string
}

func (b *BloodlinesSuite) SetupSuite() {
	httpmock.Activate()
	b.bloodlines = NewBloodlines(config.Bloodlines{
		Host: "bloodlines",
		Port: "8080",
	})
	b.url = "http://bloodlines:8080/api/"
}

func (b *BloodlinesSuite) BeforeTest() {
	httpmock.Reset()
}

func (b *BloodlinesSuite) AfterTest() {

}

func (b *BloodlinesSuite) TearDownSuite() {
	httpmock.DeactivateAndReset()
}

func TestRunBloodlinesSuite(t *testing.T) {
	s := new(BloodlinesSuite)
	suite.Run(t, s)
}

func (b *BloodlinesSuite) TestGetAllContentSuccess() {
	assert := assert.New(b.T())

	data := b.SuccessResponse()
	data.Data = []byte("[]")
	res, _ := httpmock.NewJsonResponder(200, data)

	httpmock.RegisterResponder("GET", b.url+"content?offset=0&limit=20", res)

	contents, err := b.bloodlines.GetAllContent(0, 20)

	assert.NoError(err)
	assert.NotNil(contents)
}

func (b *BloodlinesSuite) TestGetAllContentFail() {
	assert := assert.New(b.T())

	data := b.ErrorResponse("ERROR")
	res, _ := httpmock.NewJsonResponder(500, data)

	httpmock.RegisterResponder("GET", b.url+"content?offset=0&limit=20", res)

	contents, err := b.bloodlines.GetAllContent(0, 20)

	assert.Error(err)
	assert.Nil(contents)
}

func (b *BloodlinesSuite) TestNewContentSuccess() {
	assert := assert.New(b.T())

	data := b.SuccessResponse()
	content := models.NewContent("EMAIL", "text", "subject", nil)
	raw, _ := json.Marshal(content)
	data.Data = raw
	res, _ := httpmock.NewJsonResponder(200, data)

	httpmock.RegisterResponder("POST", b.url+"content", res)

	contents, err := b.bloodlines.NewContent(content)

	assert.NoError(err)
	assert.NotNil(contents)
	assert.EqualValues(content.Subject, contents.Subject)
	assert.EqualValues(content.Text, contents.Text)
	assert.EqualValues(content.Type, contents.Type)
}

func (b *BloodlinesSuite) TestNewContentFail() {
	assert := assert.New(b.T())

	data := b.ErrorResponse("ERROR")
	res, _ := httpmock.NewJsonResponder(500, data)

	httpmock.RegisterResponder("POST", b.url+"content", res)

	contents, err := b.bloodlines.NewContent(models.NewContent("EMAIL", "text", "subject", nil))

	assert.Error(err)
	assert.Nil(contents)
}

func (b *BloodlinesSuite) TestGetContentByIDSuccess() {
	assert := assert.New(b.T())

	data := b.SuccessResponse()
	c := models.NewContent("EMAIL", "text", "subject", nil)
	raw, _ := json.Marshal(c)
	data.Data = raw
	res, _ := httpmock.NewJsonResponder(200, data)

	httpmock.RegisterResponder("GET", b.url+"content/"+c.ID.String(), res)

	content, err := b.bloodlines.GetContentByID(c.ID)

	assert.NoError(err)
	assert.NotNil(content)
}

func (b *BloodlinesSuite) TestGetContentByIDError() {
	assert := assert.New(b.T())

	data := b.ErrorResponse("ERROR")
	res, _ := httpmock.NewJsonResponder(500, data)
	id := uuid.NewUUID()

	httpmock.RegisterResponder("GET", b.url+"content/"+id.String(), res)

	content, err := b.bloodlines.GetContentByID(id)

	assert.Error(err)
	assert.Nil(content)
}

func (b *BloodlinesSuite) TestUpdateContentSuccess() {
	assert := assert.New(b.T())

	data := b.SuccessResponse()
	c := models.NewContent("EMAIL", "text", "subject", nil)
	raw, _ := json.Marshal(c)
	data.Data = raw
	res, _ := httpmock.NewJsonResponder(200, data)

	httpmock.RegisterResponder("PUT", b.url+"content/"+c.ID.String(), res)

	content, err := b.bloodlines.UpdateContent(c)

	assert.NoError(err)
	assert.NotNil(content)
}

func (b *BloodlinesSuite) TestUpdateContentError() {
	assert := assert.New(b.T())

	data := b.ErrorResponse("ERROR")
	c := models.NewContent("EMAIL", "text", "subject", nil)
	raw, _ := json.Marshal(c)
	data.Data = raw
	res, _ := httpmock.NewJsonResponder(500, data)

	httpmock.RegisterResponder("PUT", b.url+"content/"+c.ID.String(), res)

	content, err := b.bloodlines.UpdateContent(c)

	assert.Error(err)
	assert.Nil(content)
}

func (b *BloodlinesSuite) TestDeleteContentSuccess() {
	assert := assert.New(b.T())

	data := b.SuccessResponse()
	res, _ := httpmock.NewJsonResponder(200, data)
	id := uuid.NewUUID()

	httpmock.RegisterResponder("DELETE", b.url+"content/"+id.String(), res)

	err := b.bloodlines.DeleteContent(id)

	assert.NoError(err)
}

func (b *BloodlinesSuite) TestDeleteContentError() {
	assert := assert.New(b.T())

	data := b.ErrorResponse("ERROR")
	res, _ := httpmock.NewJsonResponder(500, data)
	id := uuid.NewUUID()

	httpmock.RegisterResponder("DELETE", b.url+"content/"+id.String(), res)

	err := b.bloodlines.DeleteContent(id)

	assert.Error(err)
}

func (b *BloodlinesSuite) TestGetAllReceiptsSuccess() {
	assert := assert.New(b.T())

	data := b.SuccessResponse()
	data.Data = []byte("[]")
	res, _ := httpmock.NewJsonResponder(200, data)

	httpmock.RegisterResponder("GET", b.url+"receipt?offset=0&limit=20", res)

	contents, err := b.bloodlines.GetAllReceipts(0, 20)

	assert.NoError(err)
	assert.NotNil(contents)
}

func (b *BloodlinesSuite) TestGetAllReceiptsFail() {
	assert := assert.New(b.T())

	data := b.ErrorResponse("ERROR")
	res, _ := httpmock.NewJsonResponder(500, data)

	httpmock.RegisterResponder("GET", b.url+"receipt?offset=0&limit=20", res)

	contents, err := b.bloodlines.GetAllReceipts(0, 20)

	assert.Error(err)
	assert.Nil(contents)
}

func (b *BloodlinesSuite) TestSendReceiptSuccess() {
	assert := assert.New(b.T())

	data := b.SuccessResponse()
	r := models.NewReceipt(nil, uuid.NewUUID(), uuid.NewUUID())
	raw, _ := json.Marshal(r)
	data.Data = raw
	res, _ := httpmock.NewJsonResponder(200, data)

	httpmock.RegisterResponder("POST", b.url+"receipt/send", res)

	receipt, err := b.bloodlines.SendReceipt(r)

	assert.NoError(err)
	assert.NotNil(receipt)
	assert.EqualValues(receipt.ID, r.ID)
	assert.EqualValues(receipt.UserID, r.UserID)
	assert.EqualValues(receipt.ContentID, r.ContentID)
}

func (b *BloodlinesSuite) TestSendReceiptFail() {
	assert := assert.New(b.T())

	data := b.ErrorResponse("ERROR")
	res, _ := httpmock.NewJsonResponder(500, data)

	httpmock.RegisterResponder("POST", b.url+"receipt/send", res)

	receipt, err := b.bloodlines.SendReceipt(models.NewReceipt(nil, uuid.NewUUID(), uuid.NewUUID()))

	assert.Error(err)
	assert.Nil(receipt)
}

func (b *BloodlinesSuite) TestGetReceiptByIDSuccess() {
	assert := assert.New(b.T())

	data := b.SuccessResponse()
	r := models.NewReceipt(nil, uuid.NewUUID(), uuid.NewUUID())
	raw, _ := json.Marshal(r)
	data.Data = raw
	res, _ := httpmock.NewJsonResponder(200, data)

	httpmock.RegisterResponder("GET", b.url+"receipt/"+r.ID.String(), res)

	receipt, err := b.bloodlines.GetReceiptByID(r.ID)

	assert.NoError(err)
	assert.NotNil(receipt)
}

func (b *BloodlinesSuite) TestGetReceiptByIDFail() {
	assert := assert.New(b.T())

	data := b.ErrorResponse("ERROR")
	res, _ := httpmock.NewJsonResponder(500, data)

	id := uuid.NewUUID()

	httpmock.RegisterResponder("GET", b.url+"receipt/"+id.String(), res)

	receipt, err := b.bloodlines.GetReceiptByID(id)

	assert.Error(err)
	assert.Nil(receipt)
}

func (b *BloodlinesSuite) TestGetAllTriggersSuccess() {
	assert := assert.New(b.T())

	data := b.SuccessResponse()
	data.Data = []byte("[]")
	res, _ := httpmock.NewJsonResponder(200, data)

	httpmock.RegisterResponder("GET", b.url+"trigger?offset=0&limit=20", res)

	triggers, err := b.bloodlines.GetAllTriggers(0, 20)

	assert.NoError(err)
	assert.NotNil(triggers)
}

func (b *BloodlinesSuite) TestGetAllTriggersFail() {
	assert := assert.New(b.T())

	data := b.ErrorResponse("ERROR")
	res, _ := httpmock.NewJsonResponder(500, data)

	httpmock.RegisterResponder("GET", b.url+"trigger?offset=0&limit=20", res)

	triggers, err := b.bloodlines.GetAllTriggers(0, 20)

	assert.Error(err)
	assert.Nil(triggers)
}

func (b *BloodlinesSuite) TestNewTriggerSuccess() {
	assert := assert.New(b.T())

	data := b.SuccessResponse()
	t := models.NewTrigger(uuid.NewUUID(), "key", nil)
	raw, _ := json.Marshal(t)
	data.Data = raw
	res, _ := httpmock.NewJsonResponder(200, data)

	httpmock.RegisterResponder("POST", b.url+"trigger", res)

	trigger, err := b.bloodlines.NewTrigger(t)

	assert.NoError(err)
	assert.NotNil(trigger)
	assert.Equal(t.ID, trigger.ID)
	assert.EqualValues(t.Key, trigger.Key)
	assert.Equal(t.ContentID, trigger.ContentID)
}

func (b *BloodlinesSuite) TestNewTriggerFail() {
	assert := assert.New(b.T())

	data := b.ErrorResponse("ERROR")
	t := models.NewTrigger(uuid.NewUUID(), "key", nil)
	res, _ := httpmock.NewJsonResponder(500, data)

	httpmock.RegisterResponder("POST", b.url+"trigger", res)

	trigger, err := b.bloodlines.NewTrigger(t)

	assert.Error(err)
	assert.Nil(trigger)
}

func (b *BloodlinesSuite) TestGetTriggerSuccess() {
	assert := assert.New(b.T())

	data := b.SuccessResponse()
	t := models.NewTrigger(uuid.NewUUID(), "key", nil)
	raw, _ := json.Marshal(t)
	data.Data = raw
	res, _ := httpmock.NewJsonResponder(200, data)

	httpmock.RegisterResponder("GET", b.url+"trigger/"+t.Key, res)

	trigger, err := b.bloodlines.GetTriggerByKey(t.Key)

	assert.NoError(err)
	assert.NotNil(trigger)
}

func (b *BloodlinesSuite) TestGetTriggerFail() {
	assert := assert.New(b.T())

	data := b.ErrorResponse("ERROR")
	t := models.NewTrigger(uuid.NewUUID(), "key", nil)
	res, _ := httpmock.NewJsonResponder(500, data)

	httpmock.RegisterResponder("GET", b.url+"trigger/"+t.Key, res)

	trigger, err := b.bloodlines.GetTriggerByKey(t.Key)

	assert.Error(err)
	assert.Nil(trigger)
}

func (b *BloodlinesSuite) TestUpdateTriggerSuccess() {
	assert := assert.New(b.T())

	data := b.SuccessResponse()
	t := models.NewTrigger(uuid.NewUUID(), "key", nil)
	raw, _ := json.Marshal(t)
	data.Data = raw
	res, _ := httpmock.NewJsonResponder(200, data)

	httpmock.RegisterResponder("PUT", b.url+"trigger/"+t.Key, res)

	trigger, err := b.bloodlines.UpdateTrigger(t)

	assert.NoError(err)
	assert.NotNil(trigger)
}

func (b *BloodlinesSuite) TestUpdateTriggerFail() {
	assert := assert.New(b.T())

	data := b.ErrorResponse("ERROR")
	t := models.NewTrigger(uuid.NewUUID(), "key", nil)
	res, _ := httpmock.NewJsonResponder(500, data)

	httpmock.RegisterResponder("PUT", b.url+"trigger/"+t.Key, res)

	trigger, err := b.bloodlines.UpdateTrigger(t)

	assert.Error(err)
	assert.Nil(trigger)
}

func (b *BloodlinesSuite) TestDeleteTriggerSuccess() {
	assert := assert.New(b.T())

	data := b.SuccessResponse()
	res, _ := httpmock.NewJsonResponder(200, data)

	httpmock.RegisterResponder("DELETE", b.url+"trigger/key", res)

	err := b.bloodlines.DeleteTrigger("key")
	assert.NoError(err)
}

func (b *BloodlinesSuite) TestDeleteTriggerFail() {
	assert := assert.New(b.T())

	data := b.ErrorResponse("ERROR")
	res, _ := httpmock.NewJsonResponder(500, data)

	httpmock.RegisterResponder("DELETE", b.url+"trigger/key", res)

	err := b.bloodlines.DeleteTrigger("key")
	assert.Error(err)
}

func (b *BloodlinesSuite) TestActivateTriggerSuccess() {
	assert := assert.New(b.T())

	data := b.SuccessResponse()
	r := models.NewReceipt(nil, uuid.NewUUID(), uuid.NewUUID())
	raw, _ := json.Marshal(r)
	data.Data = raw
	res, _ := httpmock.NewJsonResponder(200, data)

	httpmock.RegisterResponder("POST", b.url+"trigger/key/activate", res)

	request, err := b.bloodlines.ActivateTrigger("key", &models.Receipt{})

	assert.NoError(err)
	assert.NotNil(request)
}

func (b *BloodlinesSuite) TestActivateTriggerFail() {
	assert := assert.New(b.T())

	data := b.ErrorResponse("ERROR")
	res, _ := httpmock.NewJsonResponder(500, data)

	httpmock.RegisterResponder("POST", b.url+"trigger/key/activate", res)

	request, err := b.bloodlines.ActivateTrigger("key", &models.Receipt{})

	assert.Error(err)
	assert.Nil(request)
}

func (b *BloodlinesSuite) EmptyResponse() *ServiceResponse {
	return &ServiceResponse{}
}

func (b *BloodlinesSuite) SuccessResponse() *ServiceResponse {
	r := b.EmptyResponse()
	r.Success = true
	return r
}

func (b *BloodlinesSuite) ErrorResponse(msg string) *ServiceResponse {
	r := b.EmptyResponse()
	r.Success = false
	r.Msg = msg
	return r
}
