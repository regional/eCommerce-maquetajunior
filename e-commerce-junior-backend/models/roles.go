package models

import (
	"fmt"
	"gorm/db"

	"gorm.io/gorm"
)

type Role struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	Users []User `json:"users"`
}

type Roles []Role

func MigrateRoles() {
	err := db.WithDatabaseConnection(func(database *gorm.DB) error {
		database.AutoMigrate(Role{})

		admin := Role{Id: 1, Name: "admin", Description: "Admin role"}
		seller := Role{Id: 2, Name: "seller", Description: "Seller role"}
		shooper := Role{Id: 3, Name: "shooper", Description: "Customer role"}

		for _, role := range []Role{admin, seller, shooper} {
			database.FirstOrCreate(&role, role)
		}

		return nil
	})

	if err != nil { 
		fmt.Printf("Error en la migración de roles: %v\n", err)
	}
}
