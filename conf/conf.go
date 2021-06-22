package conf

import (
	"fmt"
	"os"
	"red/model"
)

func Init() {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PWD")
	database := os.Getenv("MYSQL_RED_DB")
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	fmt.Println(config)
	model.Database(config)
}
