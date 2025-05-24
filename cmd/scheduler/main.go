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

const (
	RUN_ON = 5 * time.Second
)

func main() {
	schedulerService := schedulerService.NewService()
	ctx := context.Background()

	ticker := time.NewTicker(RUN_ON)
	defer ticker.Stop()

	err := schedulerService.Schedule(ctx)
	if err != nil {
		panic("Error in starting the scheduler service: " + err.Error())
	}

	// Then run every 5 seconds
	for range ticker.C {
		err := schedulerService.Schedule(ctx)
		if err != nil {
			fmt.Printf("Error in scheduler service: %v\n", err)
		}
	}

}
