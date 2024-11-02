package services

import "go-todos-api/models"

func CreateTodoService(title string) models.Todo {
    return models.Todo{
        Title: title,
    }
}
