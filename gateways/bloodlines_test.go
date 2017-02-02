package gateways

import (
	//"encoding/json"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ghmeier/bloodlines/config"
	//"github.com/ghmeier/bloodlines/models"
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
	b.url = "https://bloodlines:8080/api/"
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
	r.Err = fmt.Errorf(msg)
	return r
}
