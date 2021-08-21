package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"jobs-list/controllers"
	"jobs-list/driver"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/jobs", controller.GetJobs(db)).Methods("GET")
	router.HandleFunc("/jobs/{id}", controller.GetJob(db)).Methods("GET")
	router.HandleFunc("/jobs", controller.AddJob(db)).Methods("POST")
	router.HandleFunc("/jobs", controller.UpdateJob(db)).Methods("PUT")
	router.HandleFunc("/jobs/{id}", controller.RemoveJob(db)).Methods("DELETE")

	fmt.Println("Server is running at port 8000")
	log.Fatal(http.ListenAndServe(":8000",
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(router)))
}
