package router

import (
	controllers "backend/Controllers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Serve() {
	router := mux.NewRouter()

	// Company Routes
	router.HandleFunc("/company", controllers.GenerateCompany).Methods("POST")
	router.HandleFunc("/company", controllers.GetCompanies).Methods("GET")

	// User Routes
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	// Order Routes
	router.HandleFunc("/order", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/order", controllers.GetOrders).Methods("GET")
	router.HandleFunc("/order/{id}", controllers.GetOrder).Methods("GET")
	router.HandleFunc("/order/{id}", controllers.GetOrders).Methods("PUT", "PATCH")
	router.HandleFunc("/order/{id}", controllers.GetOrders).Methods("DELETE")

	options := cors.Options{
		// 		AllowedOrigins: []string{},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		Debug:          false,
	}
	handler := cors.New(options).Handler(router)
	e := godotenv.Load(".env")
	if e != nil {
		panic("unable to load env")
	}
	host := os.Getenv("HOST")
	fmt.Printf("listening at: http://localhost:%s/ \n", host)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
