package handlers

import (
	"encoding/json"
	"github.com/Schariss/opa/data"
	"github.com/Schariss/opa/repository"
	"github.com/Schariss/opa/security"
	"log"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	connection := repository.GetDatabase()
	defer repository.Closedatabase(connection)
	var user data.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		err := "Error in reading body"
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(err, w)
		return
	}
	var dbuser data.User
	persist := connection.Where("email = ?", user.Email).First(&dbuser)

	//checks if email is already register or not
	if dbuser.Email != "" {
		err := "Email already in use"
		w.Header().Set("Content-Type", "application/json")
		data.ToJSON(err,w)
		return
	}
	if persist.Error != nil {
		user.Password, err = security.GenerateHashPassword(user.Password)
		if err != nil {
			log.Fatalln("error in password hash")
		}
		//insert user details in database
		connection.Create(&user)
		log.Println("create record of user")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		data.ToJSON(user,w)
		return
	}
	//Using channels
	//c := make(chan string, 200)
	//log.Println(user.Password)
	//go security.GenerateHashPassword(&user, c)
	//log.Println(<-c)
	////insert user details in database
	////connection.Create(&user)
}