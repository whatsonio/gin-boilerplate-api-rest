package server

import (
	"app/commons/guard"
	"app/pkg/account"
	"app/pkg/health"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(guard.CorsCheck())
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	healthRouter := new(health.HealthController)

	router.GET("/readyz", healthRouter.Ready)
	router.GET("/healthz", healthRouter.Status)

	v1 := router.Group("v1")
	{
		accountGroup := v1.Group("accounts")
		{
			accountRouter := new(account.AccountController)
			accountGroup.Use(guard.JwtCheck(), guard.RolesCheck("ROLE_ADMIN, ROLE_CUSTOMER")).GET("/", accountRouter.Find)
			accountGroup.Use(guard.JwtCheck(), guard.RolesCheck("ROLE_ADMIN, ROLE_USER")).POST("/", accountRouter.Create)
			accountGroup.GET("/:id", accountRouter.FindOne)
			accountGroup.PATCH("/:id", accountRouter.Update)
			accountGroup.DELETE("/:id", accountRouter.Delete)
		}
	}

	return router

}
