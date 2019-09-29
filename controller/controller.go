package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kcwebapply/svad/domain/service"
)

const (
	REGISTER       = "/register"
	SERVICES       = "/services"
	DELETE         = "/delete"
	DELETE_SERVICE = "/delete_service"

	PROXY_X = "/svad/:endpoint"
)

var (
	serviceHostsService = service.NewSerivceImpl()
	proxyHandlerService = service.NewProxyHandlerServiceImpl()
)

func SetRouter(r *gin.Engine) *gin.Engine {
	r.POST(REGISTER, serviceHostsService.RegisterHosts)
	r.GET(SERVICES, serviceHostsService.ReturnServices)
	r.DELETE(DELETE, serviceHostsService.DeleteHosts)
	r.DELETE(DELETE_SERVICE, serviceHostsService.DeleteService)
	// proxy pattern.
	r.NoRoute(func(ctx *gin.Context) {
		proxyHandlerService.RequestHandler(ctx)
	})
	return r
}
