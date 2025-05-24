package main

import (
	db "github.com/agrawaltejas01/schedulerx/internal/database"
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

}
