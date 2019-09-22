package service

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/kcwebapply/svad/common"
	"github.com/kcwebapply/svad/domain/model"
	"github.com/kcwebapply/svad/infrastructure/http_wrapper"
	"github.com/kcwebapply/svad/infrastructure/repository"
)

type ProxyHandlerServiceImpl struct {
	serviceHostsRepository repository.ServiceHostsRepository
}

func NewProxyHandlerServiceImpl() ProxyHandlerService {
	serviceHostsRepository := repository.ServiceHostsRepositoryImpl{}
	return &ProxyHandlerServiceImpl{serviceHostsRepository: &serviceHostsRepository}
}

func (this *ProxyHandlerServiceImpl) RequestHandler(ctx *gin.Context) {
	// extract svad request info from http request..
	serviceName := ctx.GetHeader(common.SERVICE_NAME_HEADER_NAME)
	requestTypeEnum := GetRequestTypeEnum(ctx.GetHeader(common.REQUEST_TYPE_HEADER_NAME))

	// generate urlstring list
	serviceHostsEntities, err := this.serviceHostsRepository.GetHostsByServiceName(serviceName)
	if err != nil {
		common.ThrowError(err)
	}
	urlList := generateUrlList(serviceHostsEntities)

	if requestTypeEnum.StringValue == RANDOM.StringValue {
		rand.Seed(time.Now().UnixNano())
		random_number := rand.Intn(len(urlList))
		requestURL := urlList[random_number]
		path := getPath(ctx)
		response, err := doRequest(requestURL+path, ctx.ContentType(), ctx.Request)
		if err != nil {
			common.ThrowError(err)
		}

		responseJsonToUser(response, ctx)

	} else if requestTypeEnum.StringValue == ALL.StringValue {
		for _, requestURL := range urlList {
			doRequest(requestURL, ctx.ContentType(), ctx.Request)
		}
	}

}

func getBody(ctx *gin.Context) string {
	buf := make([]byte, 1024)
	num, _ := ctx.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	return reqBody
}

func getPath(ctx *gin.Context) string {
	path := ctx.Request.URL.Path
	return strings.Replace(path, "/svad", "", 1)
}

func copyResponseHeader(srcHeader http.Header, ctx *gin.Context) *gin.Context {
	for k, vs := range srcHeader {
		//ctx.Header(k, vs)
		for _, v := range vs {
			fmt.Printf("k:%s,v%s\n", k, v)
			ctx.Header(k, v)
		}
	}

	return ctx
}

func responseJsonToUser(response *http.Response, ctx *gin.Context) {
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		common.ThrowError(err)
	}
	ctx.Writer.WriteHeader(response.StatusCode)
	ctx = copyResponseHeader(response.Header, ctx)
	ctx.Render(
		response.StatusCode, render.Data{
			Data: []byte(b),
		})
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

func generateUrlList(serviceHostsEntities []model.ServiceEntity) []string {
	urlList := []string{}
	for _, e := range serviceHostsEntities {
		urlList = append(urlList, e.Host)
	}
	return urlList
}
