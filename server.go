package main

import (
	"crud-echo/entity"
	"crud-echo/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	// config.CreateConnection()
	config, err := config.NewConfig(".env")
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		config.Database.Host,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("sukses konek")

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.AutoMigrate(&entity.Club{}, &entity.Manager{}, &entity.Player{}, &entity.Country{}, &entity.Tournament{})

}
