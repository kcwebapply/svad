package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kcwebapply/svad/common"
	"github.com/kcwebapply/svad/domain/model"
	"github.com/kcwebapply/svad/domain/repository"
)

// ReturnServices func returns all services registered on svad with its hosts.
func ReturnServices(ctx *gin.Context) {
	// fetch sall erivces-hosts entity
	services, err := repository.GetAllServicesAndHosts()

	if err != nil {
		common.ThrowError(err)
	}

	// genetate services-hosts map.
	serviceMapper, err := generateServiceHostsMapper(services)
	if err != nil {
		common.ThrowError(err)
	}

	ctx.JSON(http.StatusOK, serviceMapper)

}

func generateServiceHostsMapper(serviceEntities []model.ServiceEntity) (map[string][]string, error) {
	serviceMapper := map[string][]string{}

	for _, e := range serviceEntities {
		// map exists check
		if _, exist := serviceMapper[e.ServiceName]; exist {
			hosts := append(serviceMapper[e.ServiceName], e.Host)
			serviceMapper[e.ServiceName] = hosts
		} else {
			serviceMapper[e.ServiceName] = []string{e.Host}
		}
	}
	return serviceMapper, nil
}
