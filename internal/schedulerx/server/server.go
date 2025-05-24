package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	command_server "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/server"
	job_server "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/server"
)

func RegisterCmdServer() *command_server.CmdServer {
	return command_server.NewCmdServer()
}

func ServerRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(corsMiddleware())

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	cmdServer := RegisterCmdServer()

	cmdServer.CommandRoutes(router)
	job_server.JobRoutes(router)
	return router
}
