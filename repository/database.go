package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func GetDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=root dbname=userdb port=5432 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("wrong database url")
	}
	sqldb, _ := connection.DB()

	status := "up"
	if err := sqldb.Ping(); err != nil {
		status = "down"
	}
	log.Println(status)
	return connection
}

func Closedatabase(connection *gorm.DB) {
	sqldb, _ := connection.DB()
	sqldb.Close()
}