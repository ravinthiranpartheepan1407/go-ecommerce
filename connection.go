package main

import(
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	
)

func GetDBInstance() *gorm.DB{
	dbConfig := DbConfig()
	status:= fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DbName, dbConfig.Port, dbConfig.SSLMode)
	db, err := gorm.Open(postgres.Open(status), &gorm.Config{})
	if err != nil{
		fmt.Println("Falied to connect with postgres")
	} else{
		fmt.Println("Database Connection Success")
	}

	return db
}