package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kcwebapply/svad/domain/service"
)

const (
	REGISTER = "/register"
	SERVICES = "/services/"
	PROXY    = "/svad"
)

var (
	serviceHostsService = service.NewSerivceImpl()
	proxyHandlerService = service.NewProxyHandlerServiceImpl()
)

func SetRouter(r *gin.Engine) *gin.Engine {
	r.POST(REGISTER, serviceHostsService.RegisterService)
	r.GET(SERVICES, serviceHostsService.ReturnServices)
	r.GET(PROXY, proxyHandlerService.GetHandler)
	r.POST(PROXY, proxyHandlerService.PostHandler)
	r.PUT(PROXY, proxyHandlerService.PutHandler)
	r.DELETE(PROXY, proxyHandlerService.DeleteHandler)
	return r
}
