package api

import (
	"fmt"

	"github.com/ardiost/golang-clean-web-api/api/middleware"
	"github.com/ardiost/golang-clean-web-api/api/routers"
	"github.com/ardiost/golang-clean-web-api/api/validation"
	"github.com/ardiost/golang-clean-web-api/config"
	"github.com/ardiost/golang-clean-web-api/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
	}
	r.Use(gin.Logger(), gin.Recovery(), middleware.LimitByRequest())

	RegisterSwagger(r, cfg)

	api := r.Group("/api/")
	v1 := api.Group("/v1/")
	tester := v1.Group("/tester/")
	{
		health := v1.Group("/health")
		tester_router := tester.Group("/test1")
		routers.Health(health)
		routers.TesterRouter(tester_router)

	}

	v2 := api.Group("/v2/")
	{
		tester := v2.Group("/test1")
		routers.TesterRouter(tester)

	}
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "go lang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
