package job_server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ROUTE_GROUP = "/job"
)

func JobRoutes(router *gin.Engine) *gin.Engine {
	commandRouter := router.Group(ROUTE_GROUP)

	commandRouter.POST("/create", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	return router
}
