package database

import (
	"fmt"

	"gorm.io/driver/mysql"

	"os"

	"gorm.io/gorm"
)

var (
	DbConn *gorm.DB
)
func InitDB() {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, databaseName, port)
	fmt.Println(dsn)
	DbConn, err = gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/talwar?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

}
