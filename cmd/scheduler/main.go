package main

import (
	"context"

	db "github.com/agrawaltejas01/schedulerx/internal/database"
	schedulerService "github.com/agrawaltejas01/schedulerx/internal/schedulerx/scheduler/service"
	"github.com/joho/godotenv"
)

func connectDb() {
	db.Connect()
	db.Migrate()
}

func init() {
	// Load Env file
	envVariableErr := godotenv.Load(".env")
	if envVariableErr != nil {
		panic("Error in Loading Env Variable")
	}

	connectDb()
}

func main() {
	schedulerService := schedulerService.NewService()
	ctx := context.Background()

	err := schedulerService.Schedule(ctx)
	if err != nil {
		panic("Error in starting the scheduler service: " + err.Error())
	}

	// Placeholder for main logic

}
