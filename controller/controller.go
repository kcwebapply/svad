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
	// proxy pattern.
	r.NoRoute(func(ctx *gin.Context) {
		proxyHandlerService.RequestHandler(ctx)
	})
	return r
}
