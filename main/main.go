package main

import (
	"fmt"
	"github.com/remotetodo/database"
	"github.com/remotetodo/routes"
)

func main() {
	err := database.Connect()
	if err != nil {
		//fmt.Println("database err")
		panic(err)
	}
	fmt.Println("Connected to the Database successfully")

	server := routes.Route()

	err = server.Run()
	if err != nil {
		panic(err)
	}
}
