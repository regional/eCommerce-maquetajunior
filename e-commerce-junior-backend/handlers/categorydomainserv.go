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
 * Obtiene la lista de categorias
 */
func GetCategories(rw http.ResponseWriter, r *http.Request) {
	categories := models.Categories{}

	err := db.WithDatabaseConnection(func(database *gorm.DB) error {
		if err := database.Find(&categories); err.Error != nil {
			sendError(rw, http.StatusInternalServerError)
		} else {
			for i, v := range categories {
				if err := database.Where(&models.Product{CategoryId: v.Id}).Find(&v.Products); err.Error != nil {
					sendError(rw, http.StatusInternalServerError)
				} else {
					categories[i] = v
				}
			}
		}
		return nil
	})
	if err != nil {
		sendError(rw, http.StatusInternalServerError)
	}

	sendData(rw, categories, http.StatusOK)
}

/**
 * Obtiene una Categoria por su id
 */
func GetCategory(rw http.ResponseWriter, r *http.Request) {
	if category, err := getCategoryById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, category, http.StatusOK)
	}
}

/**
 * Busca la categoria indicada por su id y devuelve error si no lo encuentra
 */
func getCategoryById(r *http.Request) (models.Category, *gorm.DB) {
	vars := mux.Vars(r)
	categoryId, _ := strconv.Atoi(vars["id"])
	category := models.Category{}

	err := db.WithDatabaseConnection(func(database *gorm.DB) error {
		if err := database.First(&category, categoryId); err.Error != nil {
			return err.Error
		} else {
			if err := database.Where(&models.Product{CategoryId: category.Id}).Find(&category.Products); err.Error != nil {
				return err.Error
			} else {
				return nil
			}
		}
	})
	if err != nil {
		return category, nil
	}

	return category, nil
}

func GetCategoryComplete(rw http.ResponseWriter, r *http.Request) {
	if category, err := getCategoryById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		if productsByCategory, erro := getProductsByCategory(category); erro != nil {
			sendError(rw, http.StatusNotFound)
		} else {
			sendData(rw, productsByCategory, http.StatusOK)
		}
	}
}

func getProductsByCategory(category models.Category) (models.ProductsByCategory, error) {
	productsByCategory := models.ProductsByCategory{
		Category: category,
	}
	products := models.Products{}

	err := db.WithDatabaseConnection(func(database *gorm.DB) error {
		if err := database.Where(&models.Product{CategoryId: category.Id}).Find(&products); err.Error != nil {
			return err.Error
		} else {
			productsByCategory.Products = products
			return nil
		}
	})
	if err != nil {
		return productsByCategory, err
	}

	return productsByCategory, nil
}

/**
 * Crea una categoria
 */
func CreateCategory(rw http.ResponseWriter, r *http.Request) {
	category := models.Category{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		err := db.WithDatabaseConnection(func(database *gorm.DB) error {
			database.Create(&category)
			return nil
		})

		if err != nil {
			sendError(rw, http.StatusInternalServerError)
		} else {
			sendData(rw, category, http.StatusCreated)
		}
	}
}

/**
 * Actualiza una categoria
 */
func UpdateCategory(rw http.ResponseWriter, r *http.Request) {
	var categoryId int64

	if old_category, err := getCategoryById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		categoryId = old_category.Id

		category := models.Category{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&category); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			category.Id = categoryId

			err := db.WithDatabaseConnection(func(database *gorm.DB) error {
				database.Save(&category)
				return nil
			})
			if err != nil {
				sendError(rw, http.StatusInternalServerError)
			}

			sendData(rw, category, http.StatusAccepted)
		}

	}
}

/**
 * Elimina una Categoria
 */
func DeleteCategory(rw http.ResponseWriter, r *http.Request) {
	if category, err := getCategoryById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		err := db.WithDatabaseConnection(func(database *gorm.DB) error {
			database.Delete(&category)
			return nil
		})
		if err != nil {
			sendError(rw, http.StatusInternalServerError)
		}

		sendData(rw, category, http.StatusOK)
	}
}
