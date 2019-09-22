package http_wrapper

import "net/http"

var (
	client = new(http.Client)
)

func GetRequest(url string, orgRequest *http.Request) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header = orgRequest.Header
	return client.Do(req)
}

func PostRequest(url, contentType string, orgRequest *http.Request) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodPost, url, orgRequest.Body)
	req.Header = orgRequest.Header
	req.Header.Set("Content-Type", contentType)
	return client.Do(req)
}

func PutRequest(url, contentType string, orgRequest *http.Request) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodPut, url, orgRequest.Body)
	req.Header = orgRequest.Header
	req.Header.Set("Content-Type", contentType)
	return client.Do(req)
}

func DeleteRequest(url, contentType string, orgRequest *http.Request) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodDelete, url, orgRequest.Body)
	req.Header = orgRequest.Header
	req.Header.Set("Content-Type", contentType)
	return client.Do(req)
}
