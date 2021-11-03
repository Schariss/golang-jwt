package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

var router *mux.Router

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

func CreateRouter() {
	router = mux.NewRouter()
}

func InitializeRoute() {
	router.HandleFunc("/signup", SignUp).Methods("POST")
	//router.HandleFunc("/signin", SignIn).Methods("POST")
}

func GetDatabase() *gorm.DB {
	//databasename := "userdb"
	//database := "postgres"
	//databasepassword := "root"
	//port := 5432
	dsn := "host=localhost user=postgres password=root dbname=userdb port=5432 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		log.Fatalln("wrong database url")
	}
	sqldb, _ := connection.DB()
	err = sqldb.Ping()
	if err != nil {
		log.Fatal("database connected")
	}
	fmt.Println("connected to database")
	return connection
}

func InitialMigration() {
	connection := GetDatabase()
	defer Closedatabase(connection)
	err := connection.AutoMigrate(User{})
	log.Println(err)
}

func Closedatabase(connection *gorm.DB) {
	sqldb, _ := connection.DB()
	sqldb.Close()
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	connection := GetDatabase()
	defer Closedatabase(connection)

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		err := "Error in reading body"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Println(user)
	var dbuser User
	connection.Where("email = ?", user.Email).First(&dbuser)

	//checks if email is already register or not
	if dbuser.Email != "" {
		err := "Email already in use"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}

	//insert user details in database
	connection.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func main(){
	CreateRouter()
	InitializeRoute()
	InitialMigration()
	fmt.Println("ended?")
	server := &http.Server{
		Addr: ":9091",
		Handler: router,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	log.New(os.Stdout, "login-api", log.LstdFlags).Fatal(server.ListenAndServe())
	fmt.Println("end")
}
