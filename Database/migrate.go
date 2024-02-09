package Database

import "go_crud/Database/Models"

func init() {
	ConnectToDatabase()
}

func Migrate() error {
	DB.AutoMigrate(&Models.Todo{})
	return nil
}
