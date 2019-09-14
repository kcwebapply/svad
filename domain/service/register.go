package service

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/kcwebapply/svad/domain/model"
	"github.com/kcwebapply/svad/domain/repository"
)

// RegisterService save service - host binding.
func RegisterService(ctx *gin.Context) {
	// get service name
	var serviceName = ctx.GetHeader("service-name")

	// requestBody
	requestBody := RegisterBody{}
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	// get urlhosts
	serviceURLList := requestBody.Hosts

	// Register Services
	for _, serviceURL := range serviceURLList {
		if _, err := url.Parse(serviceURL); err != nil {
			fmt.Println("error parsing!", serviceURL)
		}

		var entity = model.ServiceEntity{ServiceName: serviceName, Host: serviceURL}
		if err := repository.SaveHosts(entity); err != nil {
			fmt.Println(err)
		}
	}

}

type RegisterBody struct {
	Hosts []string `json:"hosts" binding:"required"`
}
