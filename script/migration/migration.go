package migration

import (
	"crud-echo/entity"
	"crud-echo/internal/config"
)

func Migrate() {

	db := config.CreateConnection()
	// config, err := config.NewConfig(".env")
	// if err != nil {
	// 	panic(err)
	// }

	// db.Migrator().CreateTable(&entity.Tournament{}, &entity.Manager{}, &entity.Club{}, &entity.Country{}, &entity.Player{})

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
	// 	config.Database.Host,
	// 	config.Database.Username,
	// 	config.Database.Password,
	// 	config.Database.Name,
	// 	config.Database.Port)
	// db.AutoMigrate()

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("sukses konek")

	db.AutoMigrate(&entity.Tournament{}, &entity.Manager{}, &entity.Club{}, &entity.Country{}, &entity.Player{})

}
