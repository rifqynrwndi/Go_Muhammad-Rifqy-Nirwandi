package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
	} `json:"name"`
}

func fetchUsers(apiURL string, ch chan<- []User, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		close(ch)
		return
	}
	defer resp.Body.Close()

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		fmt.Println("Error decoding JSON:", err)
		close(ch)
		return
	}

	ch <- users
	close(ch)
}

func printUsers(ch <-chan []User, wg *sync.WaitGroup) {
	defer wg.Done()

	for users := range ch {
		fmt.Println("Users:")
		for _, user := range users {
			fmt.Println("====")
			fmt.Printf("ID: %d\n", user.ID)
			fmt.Printf("Username: %s\n", user.Username)
			fmt.Printf("Email: %s\n", user.Email)
			fmt.Printf("First name: %s\n", user.Name.FirstName)
			fmt.Printf("Last name: %s\n", user.Name.LastName)
			fmt.Println("====")
		}
	}
}

func main() {
	apiURL := "https://fakestoreapi.com/users"
	ch := make(chan []User)
	var wg sync.WaitGroup

	wg.Add(1)
	go fetchUsers(apiURL, ch, &wg)

	wg.Add(1)
	go printUsers(ch, &wg)

	wg.Wait()
}
