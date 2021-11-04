package main

import (
	"github.com/Schariss/opa/data"
	"github.com/Schariss/opa/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var router *mux.Router

func CreateRouter() {
	router = mux.NewRouter()
}

func InitializeRoute() {
	router.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	//router.HandleFunc("/signin", SignIn).Methods("POST")
}

func main(){
	CreateRouter()
	InitializeRoute()
	data.InitialUserMigration()

	server := &http.Server{
		Addr: ":9091",
		Handler: router,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 2*time.Second,
	}

	log.Fatal(server.ListenAndServe())
	log.Println("end")
}
