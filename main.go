package main

import "fmt"

func main() {
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Println("Failed to load env")
	}
}
