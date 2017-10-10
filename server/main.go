package main

import (
	"fmt"

	"github.com/garciasa/machinedirectory/server/handler"
	"github.com/garciasa/machinedirectory/server/storage/database"
)

func main() {
	service, err := database.New("root", "root", "gotest", "mysql")
	if err != nil {
		fmt.Println("Error", err)
	}

	//service.CreateStructure()
	server := handler.New(service)
	server.Run(":5555")

}
