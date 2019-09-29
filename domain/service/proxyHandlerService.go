package service

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/kcwebapply/svad/adapter"
	"github.com/kcwebapply/svad/common"
	"github.com/kcwebapply/svad/domain/model"
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
	if !strings.HasPrefix(ctx.Request.URL.Path, "/svad/") {
		common.WriteErrorResponseOnCtx(errors.New("request path doesn't begin with '/svad/' path."), 400, ctx)
	}

	// extract svad request info from http request..
	serviceName := ctx.GetHeader(common.SERVICE_NAME_HEADER_NAME)
	requestTypeEnum := GetRequestTypeEnum(ctx.GetHeader(common.REQUEST_TYPE_HEADER_NAME))
	contentType := ctx.ContentType()
	requestObject := ctx.Request
	requestPath := getOriginalPath(requestObject.URL.Path)

	// fetch registered url list by service name.
	serviceHostsEntities, err := this.serviceHostsRepository.GetHostsByServiceName(serviceName)
	if err != nil {
		common.WriteErrorResponseOnCtx(err, 500, ctx)
	}

	if len(serviceHostsEntities) == 0 {
		common.WriteErrorResponseOnCtx(errors.New("request service doesn't registered on this server"), 404, ctx)
		return
	}
	urlList := generateUrlListFromServiceHostsEntity(serviceHostsEntities)

	//proxyRequest(urlList, requestTypeEnum, ctx)
	switch requestTypeEnum.StringValue {
	case RANDOM.StringValue:
		response, err := randomProxyRequest(urlList, requestPath, contentType, requestObject)
		if err != nil {
			common.WriteErrorResponseOnCtx(err, 500, ctx)
			return
		}
		responseToUser(response, ctx)
	case ALL.StringValue:
		responseList, _ := allProxyRequest(urlList, requestPath, contentType, requestObject)
		responseToUser(responseList[0], ctx)
	}

}

func randomProxyRequest(urlList []string, path, contentType string, request *http.Request) (*http.Response, error) {
	rand.Seed(time.Now().UnixNano())
	requestURL := urlList[rand.Intn(len(urlList))]
	return adapter.ProxyRequest(requestURL+path, contentType, request)
}

func allProxyRequest(urlList []string, path, contentType string, request *http.Request) ([]*http.Response, []error) {
	responseList := [](*http.Response){}
	errorList := []error{}
	// refactor to using async request.
	for _, requestURL := range urlList {
		response, err := adapter.ProxyRequest(requestURL+path, contentType, request)
		if err != nil {
			errorList = append(errorList, err)
		} else {
			responseList = append(responseList, response)
		}
	}
	return responseList, errorList
}

func responseToUser(response *http.Response, ctx *gin.Context) {
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		common.WriteErrorResponseOnCtx(err, 500, ctx)
		return
	}
	ctx.Writer.WriteHeader(response.StatusCode)
	ctx = copyResponseHeader(response.Header, ctx)
	ctx.Render(
		response.StatusCode, render.Data{
			Data: []byte(b),
		})
}

func generateUrlListFromServiceHostsEntity(serviceHostsEntities []model.ServiceEntity) []string {
	urlList := []string{}
	for _, e := range serviceHostsEntities {
		urlList = append(urlList, e.Host)
	}
	return urlList
}

func getOriginalPath(path string) string {
	return strings.Replace(path, "/svad", "", 1)
}

func copyResponseHeader(srcHeader http.Header, ctx *gin.Context) *gin.Context {
	for k, vs := range srcHeader {
		for _, v := range vs {
			ctx.Header(k, v)
		}
	}
	return ctx
}
