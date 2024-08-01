package database

import (
	"fmt"
	"log"

	"github.com/sarrkar/chan-ta-net/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb() error {
	var err error
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		PostgresConfig().Host,
		PostgresConfig().Port,
		PostgresConfig().User,
		PostgresConfig().Password,
		PostgresConfig().DbName,
		PostgresConfig().SslMode,
		PostgresConfig().TimeZone,
	)

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

	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func Migrate() {
	dbClient.AutoMigrate(models.Ad{})
	dbClient.AutoMigrate(models.Advertiser{})
	dbClient.AutoMigrate(models.Publisher{})

	log.Println("tables migrated")
}

func CloseDb() {
	con, _ := dbClient.DB()
	con.Close()
}
