package database

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "deepak"
	dbname   = "postgres"
)

func GetConnection() *gorm.DB {

	db := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db1, err := gorm.Open("postgres", db)

	if err != nil {
		return &gorm.DB{}
	}
	return db1

}
