package models

import (
	"fmt"

	"gorm/db"

	"gorm.io/gorm"
)

type Category struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`

	Products []Product `json:"products"`
}

type Categories []Category

func MigrateCategory() {
	shoes := Category{Id: 1, Name: "Shoes", Image: "shoes.jpg"}
	tshirts := Category{Id: 2, Name: "T-Shirts", Image: "tshirts.jpg"}
	pants := Category{Id: 3, Name: "Pants", Image: "pants.jpg"}

	err := db.WithDatabaseConnection(func(database *gorm.DB) error {
		database.AutoMigrate(Category{})
		for _, category := range []Category{shoes, tshirts, pants} {
			database.FirstOrCreate(&category, category)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error en la migración de categorías: %v\n", err)
	}
}
