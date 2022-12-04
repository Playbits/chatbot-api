package database

import (
	"fmt"
	"log"

	"exampleapp.com/m/v2/config"
	"exampleapp.com/m/v2/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}
var DB DbInstance
func Init()  {
	
    DbConfig := config.Configs.Database

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DbConfig.Host, DbConfig.Port, DbConfig.Username, DbConfig.Password, DbConfig.Database)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		 Logger: logger.Default.LogMode(logger.Info),
	})
	
    if err != nil {
        log.Fatal("Failed to connect to database. \n", err)
        panic(err)
    }

    log.Println("connected")
    db.Logger = logger.Default.LogMode(logger.Info)

    log.Println("running migrations")
    db.AutoMigrate(&models.Response{})

    DB = DbInstance{
        Db: db,
    }

}
