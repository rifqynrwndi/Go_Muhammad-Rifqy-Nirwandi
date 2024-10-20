package models

import "github.com/google/uuid"

type Food struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

var Foods = []Food{
	{ID: uuid.New().String(), Name: "Nasi Goreng", Price: 15000, Description: "Enak dan lezat"},
	{ID: uuid.New().String(), Name: "Mie Goreng", Price: 12000, Description: "Lezat dan gurih"},
}
