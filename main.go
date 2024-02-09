package main

import "go_crud/Database"

func main() {
	if err := Database.Migrate(); err != nil {
		panic(err)
	}
}
