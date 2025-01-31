package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gorm/db"
	"gorm/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

/**
 * Obtiene la lista de roles
 */
func GetRoles(rw http.ResponseWriter, r *http.Request) {
	roles := models.Roles{}
	err := db.WithDatabaseConnection(func(database *gorm.DB) error {
		if err := database.Find(&roles); err.Error != nil {
			return err.Error
		} else {
			for i, v := range roles {
				if err := database.Where(&models.User{RoleId: v.Id}).Find(&v.Users); err.Error != nil {
					return err.Error
				} else {
					roles[i] = v
				}
			}

			return nil
		}
	})

	if err != nil {
		sendError(rw, http.StatusInternalServerError)
	} else {
		sendData(rw, roles, http.StatusOK)
	}
}

/**
 * Obtiene un Rol por su id
 */
func GetRole(rw http.ResponseWriter, r *http.Request) {
	if role, err := getRoleById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, role, http.StatusOK)
	}
}

/**
 * Busca el rol indicado por su id y devuelve error si no lo encuentra
 */
func getRoleById(r *http.Request) (models.Role, *gorm.DB) {
	vars := mux.Vars(r)
	roleId, _ := strconv.Atoi(vars["id"])
	role := models.Role{}

	var err *gorm.DB
	db.WithDatabaseConnection(func(database *gorm.DB) error {
		if e := database.First(&role, roleId); e.Error != nil {
			err = e
		} else {
			if e := database.Where(&models.User{RoleId: role.Id}).Find(&role.Users); e.Error != nil {
				err = e
			}
		}

		return nil
	})

	if err != nil {
		return role, err
	} else {
		return role, nil
	}
}

func GetRoleComplete(rw http.ResponseWriter, r *http.Request) {
	if role, err := getRoleById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		if usersByRole, erro := getUsersByRole(role); erro != nil {
			sendError(rw, http.StatusNotFound)
		} else {
			sendData(rw, usersByRole, http.StatusOK)
		}
	}
}

func getUsersByRole(role models.Role) (models.UsersByRole, *gorm.DB) {
	usersByRole := models.UsersByRole{
		Role: role,
	}
	users := models.Users{}

	var err *gorm.DB
	db.WithDatabaseConnection(func(database *gorm.DB) error {
		if e := database.Where(&models.User{RoleId: role.Id}).Find(&users); e.Error != nil {
			err = e
		} else {
			usersByRole.Users = users
		}
		return nil
	})

	if err != nil {
		return usersByRole, err
	} else {
		return usersByRole, nil
	}
}

/**
 * Crea un rol
 */
func CreateRole(rw http.ResponseWriter, r *http.Request) {
	role := models.Role{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&role); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		err := db.WithDatabaseConnection(func(database *gorm.DB) error {
			database.Create(&role)
			return nil
		})

		if err != nil {
			sendError(rw, http.StatusInternalServerError)
		} else {
			sendData(rw, role, http.StatusCreated)
		}
	}
}

/**
 * Actualiza un rol
 */
func UpdateRole(rw http.ResponseWriter, r *http.Request) {
	var roleId int64

	if old_role, err := getRoleById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		roleId = old_role.Id

		role := models.Role{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&role); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			role.Id = roleId
			err := db.WithDatabaseConnection(func(database *gorm.DB) error {
				database.Save(&role)
				return nil
			})

			if err != nil {
				sendError(rw, http.StatusInternalServerError)
			} else {
				sendData(rw, role, http.StatusAccepted)
			}
		}

	}
}

/**
 * Elimina un rol
 */
func DeleteRole(rw http.ResponseWriter, r *http.Request) {
	if role, err := getRoleById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {

		err := db.WithDatabaseConnection(func(database *gorm.DB) error {
			database.Delete(&role)
			return nil
		})

		if err != nil {
			sendError(rw, http.StatusInternalServerError)
		} else {
			sendData(rw, role, http.StatusOK)
		}
	}
}
