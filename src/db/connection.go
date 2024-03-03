package mysql

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbName := os.Getenv("DBNAME")

	var err error
	var dsn = user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		return
	}
	
	log.Println("Connection Opened to Database")
}