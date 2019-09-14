package repository

import (
	"github.com/gocraft/dbr"
	"github.com/kcwebapply/svad/infrastructure/db"
)

var (
	Db   *dbr.Connection
	Sess *dbr.Session
)

func init() {
	Db = db.GetConnection()
	//Db.SetMaxOpenConns(10)
	Sess = Db.NewSession(nil)
}
