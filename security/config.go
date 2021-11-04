package security

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//Using channels
//func GenerateHashPassword(user *data.User, c chan string) {
//	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
//	user.Password = string(bytes)
//	c <- user.Password
//}