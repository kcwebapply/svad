package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kcwebapply/svad/infrastructure/repository"
)

type ProxyHandlerServiceImpl struct {
	serviceHostsRepository repository.ServiceHostsRepository
}

func NewProxyHandlerServiceImpl() ProxyHandlerService {
	serviceHostsRepository := repository.ServiceHostsRepositoryImpl{}
	return &ProxyHandlerServiceImpl{serviceHostsRepository: &serviceHostsRepository}
}

func (this *ProxyHandlerServiceImpl) GetHandler(ctx *gin.Context) {
	requestTypeEnum := GetRequestTypeEnum(ctx.GetHeader("request-type"))
	handleRequest(ctx, requestTypeEnum)
}

func (this *ProxyHandlerServiceImpl) PostHandler(ctx *gin.Context) {
	requestTypeEnum := GetRequestTypeEnum(ctx.GetHeader("request-type"))
	contentType := ctx.ContentType()
	//body := getBody(ctx)
	//http.Post(url, contentType, ctx.Request.Body)
	fmt.Println("--> type:" + requestTypeEnum.StringValue)
	fmt.Println("--> cont:" + contentType)
	fmt.Println("--> body", ctx.Request.Body)
	fmt.Println("post!!")
	//http.Post("http://localhost:8888"+ctx.Request.RequestURI, contentType, ctx.Request.Body)
	//handleRequest(ctx, requestTypeEnum)
}

func (this *ProxyHandlerServiceImpl) DeleteHandler(ctx *gin.Context) {
	requestTypeEnum := GetRequestTypeEnum(ctx.GetHeader("request-type"))
	handleRequest(ctx, requestTypeEnum)
}

func (this *ProxyHandlerServiceImpl) PutHandler(ctx *gin.Context) {
	requestTypeEnum := GetRequestTypeEnum(ctx.GetHeader("request-type"))
	handleRequest(ctx, requestTypeEnum)
}

func getBody(ctx *gin.Context) string {
	buf := make([]byte, 1024)
	num, _ := ctx.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	return reqBody
}

func handleRequest(ctx *gin.Context, requestTypeEnum *RequestTypeEnum) {
	request := ctx.Request
	if request.Method == "GET" {

	}
}
