package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	commandModel "github.com/agrawaltejas01/schedulerx/internal/schedulerx/command/model"
	jobModel "github.com/agrawaltejas01/schedulerx/internal/schedulerx/job/model"
)

var Database *gorm.DB

func Connect(config *DBConfig) {
	var err error

	dbConfig := GetDBConfig(config)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DatabaseName)

	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	db, connPoolErr := Database.DB()
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	if connPoolErr != nil {
		panic(connPoolErr)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

func Migrate() {
	Database.AutoMigrate(&commandModel.Command{})
	Database.AutoMigrate(&commandModel.Params{})
	Database.AutoMigrate(&jobModel.Job{})
}
