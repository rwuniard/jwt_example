package config

import "jwt_example/models"

func CreateDB() {
	DB.AutoMigrate(&models.User{})
}
