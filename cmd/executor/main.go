package main

import (
	"context"
	"fmt"
	"time"

	db "github.com/agrawaltejas01/schedulerx/internal/database"
	executorService "github.com/agrawaltejas01/schedulerx/internal/schedulerx/executor/service"
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

func execute(ctx context.Context, service *executorService.Service) {
	err := service.Execute(ctx)
	if err != nil {
		fmt.Println("Error in starting the scheduler service: " + err.Error())
	}
}

func main() {
	executorService := executorService.NewService()
	ctx := context.Background()

	ticker := time.NewTicker(RUN_ON)
	defer ticker.Stop()

	// First run to ensure the scheduler starts immediately
	execute(ctx, executorService)

	for range ticker.C {
		execute(ctx, executorService)
	}
}
