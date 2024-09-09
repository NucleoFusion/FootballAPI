package main

import (
	"fmt"
	"net/http"

	"api.com/example/db"
	Clubs "api.com/example/routes"
)

func main() {
	client, err := db.ConnectToDB()
	if err != nil {
		fmt.Println(err)
	}

	clubdata := db.GetCollection("clubdata", &client)

	// handling /clubs/get
	c := Clubs.ClubHandler{
		Collection: clubdata,
	}
	go http.Handle("/clubs/get/all", &c)

	//handling /clubs/get/query

	cq := Clubs.ClubQuery{
		Collection: clubdata,
	}
	go http.Handle("/clubs/get/query", &cq)

	http.ListenAndServe(":8080", nil)
}
