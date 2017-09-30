package main

import (
	"fmt"
	"github.com/garciasa/machinedirectory/server/handler"
	"github.com/garciasa/machinedirectory/server/storage/database"
)

func main() {
	service, err := database.New("root", "root", "gotest")
	if err != nil {
		fmt.Println("Error", err)
	}

	//service.CreateStructure()
	server := handler.New(service)
	server.Run(":5555")


	// test := storage.Item{
	// 	IP:         "192.168.12.1",
	// 	DomainName: "tstblab.edenirenland.com",
	// 	Deleted:    false,
	// 	Tags:       "crm, staging, pepe",
	// }

	// fmt.Println(test)

	// if err := service.Create(&test); err != nil {
	// 	fmt.Printf("Error: %s", err)
	// }

}
