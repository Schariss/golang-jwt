package data

import (
	"github.com/Schariss/opa/repository"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func InitialUserMigration() {
	connection := repository.GetDatabase()
	defer repository.Closedatabase(connection)
	if err := connection.AutoMigrate(User{}); err!= nil {
		log.Fatalln(err)
	}
}