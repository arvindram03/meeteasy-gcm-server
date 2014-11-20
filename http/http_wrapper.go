package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type HttpWrapper struct {
}

func (this *HttpWrapper) CreateRequest(url string, method string) *Request {
	request := &Request{}
	var err error
	request.request, err = http.NewRequest(method, url, nil)
	if err != nil {
		log.Println("[InternalServerError] creating request: ", url, err.Error())
	}
	return request
}

type Request struct {
	request *http.Request
}

func (this *Request) SetBasicAuth(username string, password string) {
	this.request.SetBasicAuth(username, password)
}

func (this *Request) SetHeaders(headers map[string]string) {
	for key, value := range headers {
		this.request.Header.Add(key, value)
	}
}

func (this *Request) SetJsonBody(body interface{}) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Println("[Marshal Error]: ", err)
	}
	this.request.Header.Set("Content-Type", "application/json")
	this.request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}

func (this *Request) Send() (*http.Response, error) {
	// TODO: Set Header Traceability for logging
	client := &http.Client{}
	resp, err := client.Do(this.request)
	return resp, err
}
