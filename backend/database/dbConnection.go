package database

import (
	"fmt"
	"log"
  	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
)

const (
	
	Host = "localhost"
	User = "jasurbek"
	Password = "1001"
	DbName = "wallet"
	Port = "5432"
	SslMode = "disable"
	TimeZone = "Asia/Tashkent"
)

func InitDatabase() *gorm.DB {

	dsn := fmt.Sprintf(

		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		Host, User, Password, DbName, Port, SslMode, TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		 log.Fatalf("db connection error: %v", err)
	}

	return db
}