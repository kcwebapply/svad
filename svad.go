package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/kcwebapply/svad/controller"
	"github.com/kcwebapply/svad/infrastructure/db"
	"github.com/kcwebapply/svad/infrastructure/repository"
	"github.com/kcwebapply/svad/logger"
)

func main() {

	filePath := flag.String("path", "./config.toml", "config file path")
	flag.Parse()
	db.InitDb(*filePath)
	repository.InitRepository()
	r := gin.Default()
	r.Use(logger.Logger())
	r = controller.SetRouter(r)
	r.Run(":8888")
}
