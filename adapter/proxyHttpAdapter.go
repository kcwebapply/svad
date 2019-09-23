package adapter

import (
	"fmt"
	"net/http"

	"github.com/kcwebapply/svad/infrastructure/http_wrapper"
)

func ProxyRequest(requestURL string, contentType string, request *http.Request) (*http.Response, error) {
	return doRequest(requestURL, contentType, request)

}

func doRequest(requestURL string, contentType string, request *http.Request) (*http.Response, error) {
	switch request.Method {
	case http.MethodGet:
		response, err := http_wrapper.GetRequest(requestURL, request)
		return response, err
	case http.MethodPost:
		response, err := http_wrapper.PostRequest(requestURL, contentType, request)
		return response, err
	case http.MethodPut:
		response, err := http_wrapper.PutRequest(requestURL, contentType, request)
		return response, err
	case http.MethodDelete:
		response, err := http_wrapper.DeleteRequest(requestURL, contentType, request)
		return response, err
	}
	return nil, fmt.Errorf("request method  %s doesn't supporeted on this server", request.Method)
}
