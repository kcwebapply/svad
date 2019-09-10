package service

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

func RegisterService(ctx *gin.Context) {
	serviceURL := ctx.GetHeader("service-url")
	if _, err := url.Parse(serviceURL); err != nil {
		fmt.Println("error parsing!", serviceURL)
	}

}
