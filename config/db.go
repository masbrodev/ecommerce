package config

import (
	"ecommerce/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	conn := "host=localhost user=postgres password=user dbname=mini_ecommerce port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err.Error())
	} else {
		fmt.Println("connected to database")
		DB = db
	}

	db.AutoMigrate(&entity.Product{}, &entity.Order{}, &entity.User{}, &entity.Token{})
	// db.Delete(&entity.Token{})
	db.Exec("truncate table tokens")
}
