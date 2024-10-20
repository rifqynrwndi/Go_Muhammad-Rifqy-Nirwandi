package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Response struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []User `json:"data"`
}

func main() {
	url := "https://reqres.in/api/users"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error saat membaca API: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error membaca response body: %v", err)
	}

	var responseObject Response
	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	fmt.Printf("Page: %d, Total Users: %d\n", responseObject.Page, responseObject.Total)
	for _, user := range responseObject.Data {
		fmt.Println("=========")
		fmt.Printf("ID: %d\n", user.ID)
		fmt.Printf("Name: %s %s\n", user.FirstName, user.LastName)
		fmt.Printf("Email: %s\n", user.Email)
	}
}
