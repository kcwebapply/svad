import "github.com/gin-gonic/gin"

type ServiceHostsService interface {
	RegisterService(ctx *gin.Context)
	ReturnServices(ctx *gin.Context)
}
