package services

import "my-microservice/models"

var users = []models.User{
	{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "hashed_password"},
}

func GetUsers() []models.User {
	return users
}
