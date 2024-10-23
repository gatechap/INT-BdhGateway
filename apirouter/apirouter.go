package apirouter

import (
	"net/http"

	"intbackend/bdhgateway/apicontrollers"
	"intbackend/bdhgateway/config"

	"github.com/gin-gonic/gin"

	_ "intbackend/bdhgateway/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupAPIRouter(appConfig *config.Config) *gin.Engine {

	apicontrollers.AppConfig = appConfig
	// router := gin.Default()

	router := gin.New()
	router.Use(
		//disable GIN logs from index.html
		gin.LoggerWithWriter(gin.DefaultWriter, "/index.html"),
		gin.Recovery(),
	)

	router.LoadHTMLGlob("web/*.html")

	// redis-int group: root page
	rootpage := router.Group(PATH_ROOT)
	{
		rootpage.GET(PATH_ROOT_INDEX, func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
	}

	// redis-int group: redisint
	redisint := router.Group(PATH_CONTROLLER_BDH_GATEWAY)
	{
		redisint.POST(PATH_MOETHOD_USAGECONSUMPTION, apicontrollers.UsageConsumption)
		redisint.POST(PATH_MOETHOD_GETUSAGE, apicontrollers.GetUsage)
		redisint.POST(PATH_MOETHOD_GETCREDITLIMITWITHPERFORMANCEINDICATORS, apicontrollers.GetCreditLimitWithPerformanceIndicators)
	}

	//swager
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
