package initializers

import "github.com/barbieagrawal/chapter6/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
