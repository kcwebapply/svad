package service

import (
	"net/http"

	"github.com/kcwebapply/svad/common"
)

var (
	registeredHeader = []string{"service"}
)

func copyPostRequest(url, contentType string, request *http.Request) http.Request {

	newRequest := http.Request{}

	newHeader := copyRequestHeader(request.Header)

	newRequest.Header = newHeader
	newRequest.Method = http.MethodPost
	newRequest.Body = request.Body
	newRequest.URL = request.URL

	return newRequest
}

func copyRequestHeader(srcHeader http.Header) http.Header {
	dstHeader := http.Header{}

	for k, vs := range srcHeader {

		svadHeaders := common.GetSvadHeaders()

		// skip copy header in case key is svad-related-header.
	HEADER_CHECK_LABEL:
		for _, sh := range svadHeaders {
			if k == sh {
				break HEADER_CHECK_LABEL
			}
		}

		for _, v := range vs {
			dstHeader.Add(k, v)
		}
	}

	return dstHeader
}
