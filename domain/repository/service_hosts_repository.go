package repository

import (
	"fmt"

	"github.com/gocraft/dbr"
	"github.com/kcwebapply/svad/domain/model"
	"github.com/kcwebapply/svad/infrastructure/db"
)

var dbConn *dbr.Connection
var sess *dbr.Session

const (
	tableName = "service_hosts"
)

func init() {
	dbConn = db.GetConnection()
	sess = dbConn.NewSession(nil)
	CreateTable()
}

func CreateTable() {
	_, err := sess.Exec("create table service_hosts(id serial, service_name TEXT, host TEXT, primary key(id) );")

	if err != nil {
		fmt.Println("create table error,", err)
	}
}

func GetHostsByServiceName(serviceName string) ([]model.ServiceEntity, error) {
	var serviceHosts []model.ServiceEntity
	if _, err := Sess.Select("id,service_name,host").From(tableName).Where("service_name = ?", serviceName).Load(&serviceHosts); err != nil {
		return nil, err
	}
	return serviceHosts, nil
}

func GetAllServicesAndHosts() ([]model.ServiceEntity, error) {
	var serviceHosts []model.ServiceEntity
	if _, err := Sess.Select("id,service_name,host").From(tableName).Load(&serviceHosts); err != nil {
		return nil, err
	}
	return serviceHosts, nil
}

func SaveHosts(e model.ServiceEntity) error {
	if _, err := Sess.InsertInto(tableName).Columns("service_name", "host").Record(e).Exec(); err != nil {
		return err
	}
	return nil
}
