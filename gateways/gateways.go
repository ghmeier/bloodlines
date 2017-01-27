package gateways

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ServiceResponse struct {
	Msg     string `json:"message"`
	Data    []byte `json:"data"`
	Err     error  `json:"error"`
	Success bool   `json:"success"`
}

func ServiceGet(c *http.Client, url string) ([]byte, error) {
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ServicePost(c *http.Client, url string, data []byte) ([]byte, error) {
	r := bytes.NewBuffer(data)
	resp, err := c.Post(url, "application/json", r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rData, err := handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rData, nil
}

func ServicePut(c *http.Client, url string, data []byte) ([]byte, error) {
	r := bytes.NewBuffer(data)
	req, err := http.NewRequest("PUT", url, r)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rData, err := handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rData, nil
}

func ServiceDelete(c *http.Client, url string, data []byte) ([]byte, error) {
	r := bytes.NewBuffer(data)
	req, err := http.NewRequest("DELETE", url, r)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rData, err := handleResponse(resp)
	if err != nil {
		return nil, err
	}

	return rData, nil
}

func handleResponse(resp *http.Response) ([]byte, error) {
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
		return nil, response.Err
	}

	return response.Data, nil
}
