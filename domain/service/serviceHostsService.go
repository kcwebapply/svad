package service

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/kcwebapply/svad/common"
	"github.com/kcwebapply/svad/domain/model"
	"github.com/kcwebapply/svad/infrastructure/repository"
)

type ServiceHostsServiceImpl struct {
	serviceHostsRepository repository.ServiceHostsRepository
}

func NewSerivceImpl() ServiceHostsService {
	serviceHostsRepository := repository.ServiceHostsRepositoryImpl{}
	return &ServiceHostsServiceImpl{serviceHostsRepository: &serviceHostsRepository}
}

// RegisterService save service - host binding.
func (this *ServiceHostsServiceImpl) RegisterService(ctx *gin.Context) {
	// get service name
	var serviceName = ctx.GetHeader(common.SERVICE_NAME_HEADER_NAME)
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

		urlObj, err := url.Parse(serviceURL)

		if err != nil {
			common.ThrowError(err)
		}

		urlString := generateDomainName(urlObj)

		var entity = model.ServiceEntity{ServiceName: serviceName, Host: urlString}

		if err := this.serviceHostsRepository.SaveHosts(entity); err != nil {
			common.ThrowError(err)
		}
	}
}

// ReturnServices returns all services and its hosts
func (this *ServiceHostsServiceImpl) ReturnServices(ctx *gin.Context) {
	// fetch sall erivces-hosts entity
	services, err := this.serviceHostsRepository.GetAllServicesAndHosts()

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

type RegisterBody struct {
	Hosts []string `json:"hosts" binding:"required"`
}

func generateDomainName(urlObj *url.URL) string {
	scheme := urlObj.Scheme
	suffix := "://"
	hostName := urlObj.Hostname()
	domain := scheme + suffix + hostName
	return domain
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
