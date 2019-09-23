package common

import (
	"github.com/gin-gonic/gin"
)

func WriteErrorResponseOnCtx(err error, statusCode int, ctx *gin.Context) {
	var errorResponse errorResponseStruct = errorResponseStruct{Messate: err.Error()}
	ctx.JSON(statusCode, errorResponse)
}

type errorResponseStruct struct {
	Messate string
}
