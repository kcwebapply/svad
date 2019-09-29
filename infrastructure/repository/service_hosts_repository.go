package repository

import (
	"github.com/kcwebapply/svad/domain/model"
)

const (
	tableName = "service_hosts"
)

type ServiceHostsRepositoryImpl struct {
}

func CreateServiceHostsTable() {
	sess.Exec("create table service_hosts(id serial, service_name TEXT, host TEXT, primary key(id) );")
}

func (repository *ServiceHostsRepositoryImpl) GetHostsByServiceName(serviceName string) ([]model.ServiceEntity, error) {
	var serviceHosts []model.ServiceEntity
	if _, err := sess.Select("id,service_name,host").From(tableName).Where("service_name = ?", serviceName).Load(&serviceHosts); err != nil {
		return nil, err
	}
	return serviceHosts, nil
}

func (repository *ServiceHostsRepositoryImpl) GetAllServicesAndHosts() ([]model.ServiceEntity, error) {
	var serviceHosts []model.ServiceEntity
	if _, err := sess.Select("id,service_name,host").From(tableName).Load(&serviceHosts); err != nil {
		return nil, err
	}
	return serviceHosts, nil
}

func (repository *ServiceHostsRepositoryImpl) SaveHost(e model.ServiceEntity) error {
	if _, err := sess.InsertInto(tableName).Columns("service_name", "host").Record(e).Exec(); err != nil {
		return err
	}
	return nil
}

func (repository *ServiceHostsRepositoryImpl) DeleteHost(e model.ServiceEntity) error {
	if _, err := sess.DeleteFrom(tableName).Where("service_name=?", e.ServiceName).Where("host =?", e.Host).Exec(); err != nil {
		return err
	}
	return nil
}

func (repository *ServiceHostsRepositoryImpl) DeleteService(serviceName string) error {
	if _, err := sess.DeleteFrom(tableName).Where("service_name=?", serviceName).Exec(); err != nil {
		return err
	}
	return nil
}
