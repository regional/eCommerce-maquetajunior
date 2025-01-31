package models

import (
	"fmt"

	"gorm/db"

	"gorm.io/gorm"
)

type ProductsByCategory struct {
	Category Category  `json:"category"`
	Products []Product `json:"users"`
}

type ProductsByCategories []ProductsByCategory

func MigrateProductByCategory() {
	err := db.WithDatabaseConnection(func(database *gorm.DB) error {
		database.AutoMigrate(ProductsByCategory{})
		return nil
	})
	if err != nil {
		fmt.Printf("Error en la migración de productos por categorías: %v\n", err)
	}
}
