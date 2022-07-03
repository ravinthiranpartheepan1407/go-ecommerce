package main

import "fmt"

import "github.com/joho/godotenv"


func main() {
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Println("Failed to load env")
	}

	 if GetDBInstance() != nil{
		fmt.Println("Database Connection Success")
	 }

}
