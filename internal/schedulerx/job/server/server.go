package job_server

import (
	"net/http"

	job_service "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/service"
	"github.com/gin-gonic/gin"
)

const (
	ROUTE_GROUP = "/job"
)

func JobRoutes(router *gin.Engine) *gin.Engine {

	JobService := job_service.NewService()
	commandRouter := router.Group(ROUTE_GROUP)

	commandRouter.POST("/create", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	commandRouter.GET("/get/:command", func(ctx *gin.Context) {
		job, err := JobService.GetJobsByCommand(ctx, ctx.Param("command"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"jobs": job, "status": "OK"})

	})

	return router
}
