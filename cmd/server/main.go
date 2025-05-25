package main

import (
	"context"
	"fmt"
	"os"
	"time"

	schedulerx "github.com/agrawaltejas01/schedulerx"
	db "github.com/agrawaltejas01/schedulerx/internal/database"
	"github.com/agrawaltejas01/schedulerx/internal/schedulerx/server"
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

	ctx := context.Background()

	schedulerFreq := 3 * time.Second
	executorFreq := 5 * time.Second

	schedulerx.NewSchedulerX(ctx, schedulerFreq, executorFreq)

	router := server.ServerRoutes()

	PORT := ":" + os.Getenv("PORT")

	router.Run(PORT)
	fmt.Printf("Server Running on Port: %s", PORT)
}
