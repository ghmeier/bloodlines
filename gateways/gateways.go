package gateways

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*ServiceResponse is used to unmarshal data from expresso services.
It can be used in conjunction with the following methods like this:

data, _ := ServiceGet(url)
var c []*models.Content
err := json.Unmarshal(data, &c)

Since ServiceGet returns data as a []byte, we can unmarshal it
to whatever is needed in the calling method. Here, its []*models.Content
*/
type ServiceResponse struct {
	Msg     string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success,omitempty"`
}

/*BaseService has generic methods for sending and parsing expresso service
  responses */
type BaseService struct {
	Client *http.Client
}

/*NewBaseService returns a BaseService with the defualt http Client*/
func NewBaseService() *BaseService {
	return &BaseService{
		Client: &http.Client{},
	}
}

/*ServiceSend sends a request of type METHOD to the url with data as the
  JSON payload and puts the response into i*/
func (b *BaseService) ServiceSend(method string, url string, data interface{}, i interface{}) error {
	var r *bytes.Buffer
	var err error
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return err
		}
		r = bytes.NewBuffer(b)
	} else {
		r = nil
	}

	var req *http.Request
	if r != nil {
		req, err = http.NewRequest(method, url, r)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return err
	}

	raw, err := b.doRequest(req)
	if err != nil {
		return err
	}

	if i == nil {
		return nil
	}

	err = json.Unmarshal(raw, i)
	if err != nil {
		return err
	}

	return nil
}

func (b *BaseService) doRequest(req *http.Request) ([]byte, error) {
	resp, err := b.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rData, err := b.handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rData, nil
}

func (b *BaseService) handleResponse(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response ServiceResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if !response.Success {
		if response.Msg != "" {
			return nil, fmt.Errorf("%s", response.Msg)
		} else {
			return nil, fmt.Errorf("ERROR: unknown error")
		}
	}

	return json.Marshal(response.Data)
}
