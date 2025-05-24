package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	command_server "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/server"
	job_server "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/server"
)

func Routes() *gin.Engine {
	router := gin.Default()

	router.Use(corsMiddleware())

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	command_server.CommandRoutes(router)
	job_server.JobRoutes(router)
	return router
}
