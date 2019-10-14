package api

import (
	"log"
	"net/http"
	"strings"
)

type TransportClient interface {
	Get(url string) *http.Response
	Post(url string, body string) *http.Response
	Delete(url string) *http.Response
}

type HttpTransportClient struct {
	HttpClient                      http.Client
	RetryAttemptsNetworkErrorCount  int
	RetryAttemptsInvalidStatusCount int
}

const contentTypeHeader string = "Content-Type"

func (httpTransportClient HttpTransportClient) Get(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Panic("Error while transporting get request: ", err)
	}

	return resp
}

func (httpTransportClient HttpTransportClient) Post(url string, body string) *http.Response {
	resp, err := http.Post(url, contentTypeHeader, strings.NewReader(body))
	if err != nil {
		log.Panic("Error while transporting post request: ", err)
	}

	return resp
}

func (httpTransportClient HttpTransportClient) Delete(url string) *http.Response {
	req, reqErr := http.NewRequest("DELETE", url, nil)
	if reqErr != nil {
		log.Panic("Error while creating delete request: ", reqErr)
	}

	resp, respErr := httpTransportClient.HttpClient.Do(req)
	if respErr != nil {
		log.Panic("Error while transporting delete request: ", respErr)
	}

	return resp
}
