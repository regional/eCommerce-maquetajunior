package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"gorm/db"
	"gorm/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("my_secret_key")

func GetSessionUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserByCredentials(rw, r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		token, err := generateJWT(user)
		if err != nil {
			sendError(rw, http.StatusInternalServerError)
		} else {
			sendData(rw, token, http.StatusOK)
		}
	}
}

func getUserByCredentials(rw http.ResponseWriter, r *http.Request) (models.User, *gorm.DB) {
	// Crear el objeto vacio
	user := models.User{}

	// objeto para recibir usuario y contraseña
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&credentials); err != nil {
		http.Error(rw, "Invalid request payload", http.StatusBadRequest)
		return user, nil
	}

	var err *gorm.DB
	db.WithDatabaseConnection(func(database *gorm.DB) error {
		if e := database.Preload("Role").Where("username = ?", credentials.Username).First(&user); e.Error != nil {
			http.Error(rw, "Usuario no encontrado", http.StatusBadRequest)
			err = e
		}
		return nil
	})
	if err != nil {
		http.Error(rw, "Usuario no encontrado", http.StatusBadRequest)
	}

	// Compare the hashed password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		http.Error(rw, "La contraseña es incorrecta", http.StatusUnauthorized)
		return user, nil
	}

	return user, nil
}

func generateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"roleid":   user.RoleId,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"avatar":   user.Avatar,
		"rolename": user.Role.Name,
		"userid":   user.Id,
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
