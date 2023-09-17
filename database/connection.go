package database

import (
	"fmt"

	"github.com/mrojasb2000/go-jwt-authentication-app/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	username := helpers.GetEnvVariable("USERNAME")
	password := helpers.GetEnvVariable("PASSWORD")
	dbname := helpers.GetEnvVariable("DBNAME")

	connectionString := fmt.Sprintf("%s:%s@/%s", username, password, dbname)
	_, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database.")
	}
}
