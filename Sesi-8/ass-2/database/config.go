package database

import (
	"ass-2/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dsn := "host=localhost user=postgres password=12345678 dbname=orders_by sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed connect to database")
	}

	err = db.AutoMigrate(&models.Orders{}, &models.Items{})

	if err != nil {
		errMsg := fmt.Sprintf("Failed to migrate models: %s", err.Error())
		panic(errMsg)
	}

	fmt.Println("Database Connected")

	return db
}
