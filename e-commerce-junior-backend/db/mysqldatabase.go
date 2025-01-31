package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database es la conexión a la base de datos
var Database = func() (db *gorm.DB) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error en la conexión", err)
		panic(err)
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println("Error al obtener la conexión subyacente: %v", err)
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)

		fmt.Println("Conexión exitosa")
		return db
	}
}


func WithDatabaseConnection(operation func(*gorm.DB) error) error {
    database := Database()
    sqlDB, err := database.DB()
    if err != nil {
        return fmt.Errorf("error al obtener la conexión subyacente: %v", err)
    }
    defer func() {
        if err := sqlDB.Close(); err != nil {
            fmt.Printf("error al cerrar la conexión: %v\n", err)
        }
    }()

    return operation(database)
}