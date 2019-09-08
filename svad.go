package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kcwebapply/svad/controller"
	"github.com/kcwebapply/svad/logger"
)

func main() {
	r := gin.Default()
	r.Use(logger.Logger())
	r = controller.SetRouter(r)
	r.Run(":8888")
}
