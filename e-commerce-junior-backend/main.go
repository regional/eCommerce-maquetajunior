package main

import (
	"log"
	"net/http"

	"gorm/handlers"
	"gorm/middleware"
	"gorm/models"

	"github.com/gorilla/mux"
)

func main() {


	models.MigrateRoles()
	models.MigrateCategory()
	models.MigrateProduct()
	models.MigrateUser()

	mux := mux.NewRouter()

	// Roles Service
	mux.Handle("/api/role/", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetRoles), "admin")).Methods(http.MethodGet)
	mux.Handle("/api/role/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetRole), "admin")).Methods(http.MethodGet)
	mux.Handle("/api/role/userByRole/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetRoleComplete), "admin")).Methods(http.MethodGet)
	mux.Handle("/api/role/", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.CreateRole), "admin")).Methods(http.MethodPost)
	mux.Handle("/api/role/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.UpdateRole), "admin")).Methods(http.MethodPut)
	mux.Handle("/api/role/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.DeleteRole), "admin")).Methods(http.MethodDelete)

	// User Service
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods(http.MethodPost)
	mux.Handle("/api/user/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetUser), "admin","seller","shooper")).Methods(http.MethodGet)
	mux.Handle("/api/user/", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetUsers),"admin")).Methods(http.MethodGet)
	mux.Handle("/api/user/userByRole/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetUsersByRole),"admin")).Methods(http.MethodGet)
	mux.Handle("/api/user/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.UpdateUser),"admin")).Methods(http.MethodPut)
	mux.Handle("/api/user/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.DeleteUser),"admin")).Methods(http.MethodDelete)

	// Category Service
	mux.Handle("/api/category/", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetCategories), "admin","seller")).Methods(http.MethodGet)
	mux.Handle("/api/category/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetCategory), "admin","seller","shooper")).Methods(http.MethodGet)
	mux.Handle("/api/category/productsByCategory/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetCategoryComplete),"admin","seller","shooper")).Methods(http.MethodGet)
	mux.Handle("/api/category/", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.CreateCategory),"admin","seller")).Methods(http.MethodPost)
	mux.Handle("/api/category/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.UpdateCategory),"admin","seller")).Methods(http.MethodPut)
	mux.Handle("/api/category/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.DeleteCategory),"admin")).Methods(http.MethodDelete)

	// Product Service
	mux.Handle("/api/product/", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetProducts),"admin","seller","shooper")).Methods(http.MethodGet)
	mux.Handle("/api/product/productByCategory/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetProductsByCategory),"admin","seller","shooper")).Methods(http.MethodGet)
	mux.Handle("/api/product/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetProduct),"admin","seller","shooper")).Methods(http.MethodGet)
	mux.Handle("/api/product/", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.CreateProduct),"admin","seller")).Methods(http.MethodPost)
	mux.Handle("/api/product/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.UpdateProduct),"admin","seller")).Methods(http.MethodPut)
	mux.Handle("/api/product/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.DeleteProduct),"admin","seller")).Methods(http.MethodDelete)
	
	// ShopingCar Service
	mux.Handle("/api/shoppingcart/", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetShopingCar),"admin","seller","shooper")).Methods(http.MethodGet)
	mux.Handle("/api/shoppingcart/", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.SaveShopingCar),"admin","seller", "shooper")).Methods(http.MethodPost)
	// Chat service
	mux.Handle("/api/chatbot/", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.GetChatMessagesHandler),"admin","seller", "shooper")).Methods(http.MethodGet)
	mux.Handle("/api/chatbot/", middleware.JWTAuthMiddleware(http.HandlerFunc(handlers.CreateChatMessageHandler),"admin","seller", "shooper")).Methods(http.MethodPost)

	// Sesion Service
	mux.HandleFunc("/api/session/", handlers.GetSessionUser).Methods(http.MethodPost)
	mux.HandleFunc("/api/healt/", handlers.Healt).Methods(http.MethodGet)
	
	//Aplica el middleware de CORS
	wrappedMux := middleware.EnableCORS(mux)
	log.Fatal(http.ListenAndServe(":3000", wrappedMux))
}
