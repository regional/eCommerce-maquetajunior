package models

import (
	"errors"
	"fmt"

	"gorm/db"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleId   int64  `json:"roleId"`
	Avatar   string `json:"avatar"`

	Role Role `json:"role" gorm:"foreignKey:RoleId;references:Id"`
}

type Users []User

func MigrateUser() {
	err := db.WithDatabaseConnection(func(database *gorm.DB) error {
		database.AutoMigrate(User{})

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin1234"), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("error al hashear la contraseña: %v", err)
		}

		password := string(hashedPassword)

		admin := User{Id: 1, Username: "admin", Password: password, Email: "admin@gmail.com", RoleId: 1, Avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSzBpnouxDuF063trW5gZOyXtyuQaExCQVMYA&s"}

		result := database.Where("id = ?", admin.Id).First(&admin)
        if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
            database.Create(&admin)
        }

        return nil
	})
	if err != nil {
		fmt.Printf("Error en la migración del usuario: %v\n", err)
	}
}
