package repository

import (
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
	//dbConn.SetMaxOpenConns(10)
	sess = dbConn.NewSession(nil)
}

func CreateTable() {
	_, _ = sess.Exec("create table service_hosts(id INTEGER PRIMARY KEY AUTOINCREMENT, service_name TEXT, host TEXT);")
}

func GetHostsByServiceName(serviceName string) ([]model.ServiceEntity, error) {
	var serviceHosts []model.ServiceEntity
	if _, err := Sess.Select("id,service_name,host").From(tableName).Where("service_name = ?", serviceName).Load(&serviceHosts); err != nil {
		return nil, err
	}
	return serviceHosts, nil
}

func SaveHosts(e model.ServiceEntity) error {
	if _, err := Sess.InsertInto(tableName).Columns("id", "service_name", "host").Record(e).Exec(); err != nil {
		return err
	}
	return nil
}
