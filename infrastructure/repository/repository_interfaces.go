package repository

import "github.com/kcwebapply/svad/domain/model"

// interface to fetch data from service_hosts table.
type ServiceHostsRepository interface {
	GetHostsByServiceName(string) ([]model.ServiceEntity, error)
	GetAllServicesAndHosts() ([]model.ServiceEntity, error)
	SaveHosts(model.ServiceEntity) error
}
