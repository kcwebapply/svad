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

func SetRouter(r *gin.Engine) *gin.Engine {
	r.POST(REGISTER, service.RegisterService)
	r.GET(SERVICES, service.ReturnServices)
	r.GET(PROXY, service.GetHandler)
	r.POST(PROXY, service.PostHandler)
	r.PUT(PROXY, service.PutHandler)
	r.DELETE(PROXY, service.DeleteHandler)
	return r
}
