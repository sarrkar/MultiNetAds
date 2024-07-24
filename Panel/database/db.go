package database

import (
	"fmt"
	"log"

	"github.com/sarrkar/Chan-ta-net/Panel/config"
	"github.com/sarrkar/Chan-ta-net/Panel/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb() error {
	var err error
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tehran",
		config.Config().Postgres.Host, config.Config().Postgres.Port, config.Config().Postgres.User, config.Config().Postgres.Password,
		config.Config().Postgres.DbName)

	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}
	log.Println("Db connection established")

	dbClient.AutoMigrate(&models.Ad{})
	dbClient.AutoMigrate(&models.Advertiser{})
	dbClient.AutoMigrate(&models.Publisher{})
	log.Println("tables migrated")

	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	con, _ := dbClient.DB()
	con.Close()
}
