package config

import (
	"rest-api-golang/structs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dsn := "host=localhost user=postgres password=12345678 dbname=db-go-sql sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
