package repository

import (
	"github.com/gocraft/dbr"
	"github.com/kcwebapply/svad/infrastructure/db"
)

var (
	dbConn *dbr.Connection
	sess   *dbr.Session
)

func init() {
	dbConn = db.GetConnection()
	//Db.SetMaxOpenConns(10)
	sess = dbConn.NewSession(nil)
	CreateServiceHostsTable()
}
