package command_server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	commandModel "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/model"
	commandService "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/service"
)

func CommandRoutes(router *gin.Engine) *gin.Engine {
	commandRouter := router.Group("/command")

	commandRouter.POST("/create", func(ctx *gin.Context) {
		var command commandModel.Command
		if err := ctx.ShouldBindJSON(&command); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := commandService.CreateCommand(ctx, command)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})

	})

	commandRouter.GET("/:command", func(ctx *gin.Context) {
		cmd := ctx.Param("command")

		command, err := commandService.GetCommand(ctx, cmd)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if command.Command == "" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Command not found"})
			return
		}
		fmt.Println("Command retrieved:", command)
		ctx.JSON(http.StatusOK, gin.H{"command": command, "status": "OK"})
	})

	return router
}
