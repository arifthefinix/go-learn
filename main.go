package main

import (
	"go_crud/Database"
	"go_crud/Routes"
)

func main() {
	if err := Database.Migrate(); err != nil {
		panic(err)
	}
	router := Routes.SetupRouter()
	router.Run(":8031")
}
