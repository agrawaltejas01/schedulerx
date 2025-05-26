package main

import (
	"context"
	"fmt"
	"time"

	db "github.com/agrawaltejas01/schedulerx/internal/database"
	schedulerService "github.com/agrawaltejas01/schedulerx/internal/schedulerx/scheduler/service"
	"github.com/joho/godotenv"
)

func connectDb() {
	db.Connect(nil)
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

const (
	RUN_ON = 5 * time.Second
)

func schedule(ctx context.Context, schedulerService *schedulerService.Service) {
	err := schedulerService.Schedule(ctx)
	if err != nil {
		fmt.Println("Error in starting the scheduler service: " + err.Error())
	}
}

func main() {
	schedulerService := schedulerService.NewService()
	ctx := context.Background()

	ticker := time.NewTicker(RUN_ON)
	defer ticker.Stop()

	// First run to ensure the scheduler starts immediately
	schedule(ctx, schedulerService)

	for range ticker.C {
		schedule(ctx, schedulerService)
	}
}
