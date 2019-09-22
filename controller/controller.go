package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kcwebapply/svad/domain/service"
)

const (
	REGISTER = "/register"
	SERVICES = "/services/"

	PROXY_X = "/svad/:endpoint"
)

var (
	serviceHostsService = service.NewSerivceImpl()
	proxyHandlerService = service.NewProxyHandlerServiceImpl()
)

func SetRouter(r *gin.Engine) *gin.Engine {
	r.POST(REGISTER, serviceHostsService.RegisterService)
	r.GET(SERVICES, serviceHostsService.ReturnServices)
	r.GET(PROXY_X, proxyHandlerService.RequestHandler)
	r.POST(PROXY_X, proxyHandlerService.RequestHandler)
	r.PUT(PROXY_X, proxyHandlerService.RequestHandler)
	r.DELETE(PROXY_X, proxyHandlerService.RequestHandler)
	return r
}
