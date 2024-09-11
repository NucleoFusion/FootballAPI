package main

import (
	"fmt"
	"net/http"

	"api.com/example/db"
	Clubs "api.com/example/routes/clubs"
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

	//handling /clubs/get/all/limit
	cal := Clubs.ClubAllLimit{
		Collection: clubdata,
	}
	go http.Handle("/clubs/get/all/{limit}", &cal)

	//handling /clubs/get/query/limit
	cql := Clubs.ClubAllLimit{
		Collection: clubdata,
	}
	go http.Handle("/clubs/get/query/{limit}", &cql)

	//handling /clubs/get/sortBy/{val}
	cas := Clubs.ClubAllSort{
		Collection: clubdata,
	}
	go http.Handle("/clubs/get/sortBy/{sortVal}", &cas)

	//handling /clubs/get/sortBy/{val}/{limit}
	cls := Clubs.ClubLimitSort{
		Collection: clubdata,
	}
	go http.Handle("/clubs/get/sortBy/{sortVal}/limit/{limit}", &cls)

	http.ListenAndServe(":8080", nil)
}
