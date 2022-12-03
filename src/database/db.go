package database

import (
	"flag"
	"fmt"
	"log"
	"os"

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

	configFlag := flag.String("config", "dev", "config flag can be dev for develop or prod for production")
	prodConfigPath := flag.String("config-path", "", "config-path production config file path")
	config.Init(configFlag, prodConfigPath)
    DbConfig := config.Configs.Database

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DbConfig.Host, DbConfig.Port, DbConfig.Username, DbConfig.Password, DbConfig.Database)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		 Logger: logger.Default.LogMode(logger.Info),
	})
	
    if err != nil {
        log.Fatal("Failed to connect to database. \n", err)
        os.Exit(2)
    }

    log.Println("connected")
    db.Logger = logger.Default.LogMode(logger.Info)

    log.Println("running migrations")
    db.AutoMigrate(&models.Response{})

    DB = DbInstance{
        Db: db,
    }

}
// func Init(url string) *gorm.DB {

// 	configFlag := flag.String("config", "dev", "config flag can be dev for develop or prod for production")
// 	prodConfigPath := flag.String("config-path", "", "config-path production config file path")
// 	// init service configs
// 	config.Init(configFlag, prodConfigPath)
//     DbConfig := config.Configs.Database

//     dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DbConfig.Host, DbConfig.Port, DbConfig.Username, DbConfig.Password, DbConfig.Database)
//     db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

//     if err != nil {
//         log.Fatalln(err)
//     }

//     // db.AutoMigrate(&models.Chatbot{})

//     return db
// }
