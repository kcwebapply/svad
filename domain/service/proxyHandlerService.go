package service

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kcwebapply/svad/common"
	"github.com/kcwebapply/svad/infrastructure/repository"
)

type ProxyHandlerServiceImpl struct {
	serviceHostsRepository repository.ServiceHostsRepository
}

func NewProxyHandlerServiceImpl() ProxyHandlerService {
	serviceHostsRepository := repository.ServiceHostsRepositoryImpl{}
	return &ProxyHandlerServiceImpl{serviceHostsRepository: &serviceHostsRepository}
}

func (this *ProxyHandlerServiceImpl) GetHandler(ctx *gin.Context) {
	/*requestTypeEnum := GetRequestTypeEnum(ctx.GetHeader(common.REQUEST_TYPE_HEADER_NAME))
	handleRequest(ctx, requestTypeEnum)*/
}

func (this *ProxyHandlerServiceImpl) PostHandler(ctx *gin.Context) {
	serviceName := ctx.GetHeader(common.SERVICE_NAME_HEADER_NAME)
	requestTypeEnum := GetRequestTypeEnum(ctx.GetHeader(common.REQUEST_TYPE_HEADER_NAME))

	serviceHostsEntities, err := this.serviceHostsRepository.GetHostsByServiceName(serviceName)

	if err != nil {
		common.ThrowError(err)
	}

	urlList := []string{}
	for _, e := range serviceHostsEntities {
		urlList = append(urlList, e.Host)
	}

	handleRequest(urlList, ctx, requestTypeEnum)

	//body := getBody(ctx)
	//http.Post(url, contentType, ctx.Request.Body)

	//http.Post("http://localhost:8888"+ctx.Request.RequestURI, contentType, ctx.Request.Body)
	//handleRequest(ctx, requestTypeEnum)
}

func (this *ProxyHandlerServiceImpl) DeleteHandler(ctx *gin.Context) {
	/*requestTypeEnum := GetRequestTypeEnum(ctx.GetHeader(common.REQUEST_TYPE_HEADER_NAME))
	handleRequest(ctx, requestTypeEnum)*/
}

func (this *ProxyHandlerServiceImpl) PutHandler(ctx *gin.Context) {
	/*requestTypeEnum := GetRequestTypeEnum(ctx.GetHeader(common.REQUEST_TYPE_HEADER_NAME))
	handleRequest(ctx, requestTypeEnum)*/
}

func getBody(ctx *gin.Context) string {
	buf := make([]byte, 1024)
	num, _ := ctx.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	return reqBody
}

func handleRequest(urlList []string, ctx *gin.Context, requestTypeEnum *RequestTypeEnum) {
	request := ctx.Request
	fmt.Println("size:", len(urlList))
	contentType := ctx.ContentType()
	if requestTypeEnum.StringValue == RANDOM.StringValue {
		rand.Seed(time.Now().UnixNano())
		random_number := rand.Intn(len(urlList))
		requestURL := urlList[random_number]
		doRequest(requestURL, contentType, request)
	} else if requestTypeEnum.StringValue == ALL.StringValue {
		for _, requestURL := range urlList {
			doRequest(requestURL, contentType, request)
		}
	}

}

func doRequest(requestURL string, contentType string, request *http.Request) {
	if request.Method == http.MethodPost {
		fmt.Println("post!")
		response, err := http.Post(requestURL, contentType, request.Body)
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("res:", response.Body)
	}
}
