package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Users struct {
	ID 			int		`json:"id"`
	UID			string	`json:"uid"`
	FirstName 	string	`json:"first_name"`
	LastName 	string	`json:"last_name"`
	Username 	string	`json:"username"`
	Address 	Address `json:"address"`
}

type Address struct {
	City			string		`json:"city"`
	StreetName 		string		`json:"street_name"`
	StreetAddress	string		`json:"street_address"`
	ZipCode			string		`json:"zip_code"`
	State 			string		`json:"state"`
	Country 		string		`json:"country"`
	Coordinates		Coordinates	`json:"coordinates"`
}

type Coordinates struct {
	Lat	float64 `json:"lat"`
	Lng	float64 `json:"lng"`
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/users", getUsers)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		res, err := http.Get("https://random-data-api.com/api/users/random_user?size=10")

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(res.Body)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()

		// sb := string(body)

		var users []Users

		errUnmars := json.Unmarshal(body, &users)

		if errUnmars != nil {
			fmt.Printf("Error Unmarshal %s", errUnmars)
		}

		jsonData, errMars := json.Marshal(users)

		if errMars != nil {
			fmt.Printf("Error Marshal %s", errMars)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jsonData))
	}
}