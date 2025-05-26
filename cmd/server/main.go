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
	db.Connect(nil)
	db.Migrate()
}

func init() {
	// Load Env file
	envVariableErr := godotenv.Load(".env")
	if envVariableErr != nil {
		panic("Error in Loading Env Variable")
	}

	// connectDb()
}

func main() {

	ctx := context.Background()

	schedulerFreq := 3 * time.Second
	executorFreq := 5 * time.Second

	schedulerx.NewSchedulerX(ctx, schedulerFreq, executorFreq, &schedulerx.DBConfig{
		Config: db.DBConfig{
			Host:         os.Getenv("DB_HOST"),
			Port:         os.Getenv("DB_PORT"),
			Username:     os.Getenv("DB_USERNAME"),
			Password:     os.Getenv("DB_PASSWORD"),
			DatabaseName: os.Getenv("DB_NAME"),
		},
	})

	router := server.ServerRoutes()

	PORT := ":" + os.Getenv("PORT")

	router.Run(PORT)
	fmt.Printf("Server Running on Port: %s", PORT)
}
