package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kcwebapply/svad/domain/service"
)

const (
	SVAD  = "/register/"
	PROXY = "/svad/"
)

func SetRouter(r *gin.Engine) *gin.Engine {
	r.POST(SVAD, service.RegisterService)
	r.GET(PROXY, service.GetHandler)
	r.POST(PROXY, service.PostHandler)
	r.PUT(PROXY, service.PutHandler)
	r.DELETE(PROXY, service.DeleteHandler)
	return r
}
