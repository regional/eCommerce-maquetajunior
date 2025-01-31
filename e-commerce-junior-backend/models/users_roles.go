package models

import (
	"fmt"
	"gorm/db"

	"gorm.io/gorm"
)

type UsersByRole struct {
	Role Role `json:"role"`
	Users []User `json:"users"`
}

type UsersByRoles []UsersByRole

func MigrateUserByRole(){
	err := db.WithDatabaseConnection(func(database *gorm.DB) error {
		database.AutoMigrate(UsersByRole{})
		return nil
	})
	if err != nil {
		fmt.Printf("Error en la migración de usuarios por roles: %v\n", err)
	}
}