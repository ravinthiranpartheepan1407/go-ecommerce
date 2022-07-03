package main

import "fmt"

import "github.com/joho/godotenv"

import "github.com/ravinthiranpartheepan1407/go-ecommerce"


func main() {
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Println("Failed to load env")
	}

	connection.GetDB()
}
