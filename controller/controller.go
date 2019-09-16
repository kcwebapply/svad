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
)

func SetRouter(r *gin.Engine) *gin.Engine {
	r.POST(REGISTER, serviceHostsService.RegisterService)
	r.GET(SERVICES, serviceHostsService.ReturnServices)
	r.GET(PROXY, service.GetHandler)
	r.POST(PROXY, service.PostHandler)
	r.PUT(PROXY, service.PutHandler)
	r.DELETE(PROXY, service.DeleteHandler)
	return r
}
