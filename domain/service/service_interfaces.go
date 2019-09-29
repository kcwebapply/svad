package service

import "github.com/gin-gonic/gin"

//
type ProxyHandlerService interface {
	/*	GetHandler(ctx *gin.Context)
		PostHandler(ctx *gin.Context)
		DeleteHandler(ctx *gin.Context)
		PutHandler(ctx *gin.Context)*/
	RequestHandler(ctx *gin.Context)
}

// ServiceHostsService deal with service-host entity data.
type ServiceHostsService interface {
	RegisterHosts(ctx *gin.Context)
	ReturnServices(ctx *gin.Context)
	DeleteHosts(ctx *gin.Context)
	DeleteService(ctx *gin.Context)
}
