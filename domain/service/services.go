package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kcwebapply/svad/domain/model"
	"github.com/kcwebapply/svad/domain/repository"
)

func ReturnServices(ctx *gin.Context) {
	services, err := repository.GetAllServicesAndHosts()
	if err != nil {
		fmt.Println("err!", err)
	}

	serviceMapper, err := generateServiceHostsJsonMapper(services)
	ctx.JSON(http.StatusOK, serviceMapper)

}

func generateServiceHostsJsonMapper(serviceEntities []model.ServiceEntity) (map[string]HostStruct, error) {
	serviceMapper := map[string]HostStruct{}

	for _, e := range serviceEntities {
		// map exists check
		if _, exist := serviceMapper[e.ServiceName]; exist {
			hosts := serviceMapper[e.ServiceName].Hosts
			hosts = append(hosts, e.Host)
			serviceMapper[e.ServiceName] = HostStruct{Hosts: hosts}
		} else {
			serviceMapper[e.ServiceName] = HostStruct{Hosts: []string{e.Host}}
		}
	}
	return serviceMapper, nil
}

type HostStruct struct {
	Hosts []string `json:"hosts" binding:"required"`
}
