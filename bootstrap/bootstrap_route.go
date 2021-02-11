package bootstrap

import (
	"net/http"

	"github.com/rest_api/constants"
	"github.com/rest_api/docs"
	"github.com/rest_api/handlers/rest"
	"github.com/rest_api/handlers/rest/middlewares"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func initREST() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"name":    constants.ServiceName,
			"version": constants.ServiceVersion,
		})
	})

	// Swagger
	docs.SwaggerInfo.Title = constants.ServiceName
	docs.SwaggerInfo.Version = constants.ServiceVersion
	if cfg.App.Env != "production" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	router.Use(middlewares.RequestLoggerMiddleware(), middlewares.CORSMiddleware())

	// handlers
	customerHandler := rest.NewCustomerHandler(customerService)

	v1Router := router.Group("/v1")
	{
		customerRouter := v1Router.Group("/customer")
		customerRouter.GET("/synchronize", customerHandler.Synchronize)
	}

	return router
}
