package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gorm/db"
	"gorm/models"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/**
 * Obtiene la lista de todos los registros
 */
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := models.Users{}
	err := db.WithDatabaseConnection(func(database *gorm.DB) error {
		if errr := database.Preload("Role").Find(&users).Error; errr != nil {
			sendError(rw, http.StatusNotFound)
			return errr
		} else {
			sendData(rw, users, http.StatusOK)
			return nil
		}
	})
	if err != nil {
		sendError(rw, http.StatusInternalServerError)
	}
}

func GetUsersByRole(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleId, _ := strconv.Atoi(vars["id"])
	users := models.Users{}

	err := db.WithDatabaseConnection(func(database *gorm.DB) error {
		if errr := database.Where(models.User{RoleId: int64(roleId)}).Find(&users); errr.Error != nil {
			sendError(rw, http.StatusNotFound)
			return errr.Error
		} else {
			sendData(rw, users, http.StatusOK)
			return nil
		}
	})
	if err != nil {
		sendError(rw, http.StatusInternalServerError)
	}
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getuserBYId(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, user, http.StatusOK)
	}
}

func getuserBYId(r *http.Request) (models.User, *gorm.DB) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user := models.User{}

	errr := db.WithDatabaseConnection(func(database *gorm.DB) error {
		err := database.First(&user, userId).Error
		return err
	})
	
	if errr != nil {
		return user, nil
	}

	return user, nil
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	// Crear el objeto vacio
	user := models.User{}
	// Obtiener el body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
		return
	}

	user.Password = string(hashedPassword)

	err = db.WithDatabaseConnection(func(database *gorm.DB) error {
		if errr := database.Create(&user).Error; errr != nil {
			http.Error(rw, "Error creating user", http.StatusInternalServerError)
			return errr
		}

		sendData(rw, user, http.StatusCreated)
		return nil
	})
	if err != nil {
		sendError(rw, http.StatusInternalServerError)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var userId int64

	if old_user, err := getuserBYId(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {

		userId = old_user.Id

		user := models.User{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&user); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			user.Id = userId
			// Hash the password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				http.Error(rw, "Error hashing password", http.StatusInternalServerError)
				return
			}
			user.Password = string(hashedPassword)

			err = db.WithDatabaseConnection(func(database *gorm.DB) error {
				if errr := database.Save(&user).Error; errr != nil {
					return errr
				}
				sendData(rw, user, http.StatusOK)
				return nil
			})
			if err != nil {
				sendError(rw, http.StatusInternalServerError)
			}
		}
	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	user, err := getuserBYId(r)
	if err != nil {
		sendError(rw, http.StatusNotFound)
		return
	}

	errr := db.WithDatabaseConnection(func(database *gorm.DB) error {
		if err := database.Delete(&user).Error; err != nil {		
			return err
		} else {
			return nil
		}
	})

	if errr != nil {
		sendError(rw, http.StatusInternalServerError)
	} else {
		sendData(rw, user, http.StatusOK)
	}
}
