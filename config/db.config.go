package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	database *gorm.DB
	err      error
)

func DatabaseInit() {
	/*
		You could change the database config 
		depends on your database setup.
	*/
	host := "localhost"
	user := "postgres" 
	password := "15"
	dbName := "crud"
	port := 5432

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbName, port)
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbl_",
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return database
}
