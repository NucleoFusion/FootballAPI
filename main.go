package main

import (
	"fmt"
	"net/http"

	"api.com/example/db"
	Clubs "api.com/example/routes/clubs"
	Stadiums "api.com/example/routes/stadiums"
)

func main() {
	client, err := db.ConnectToDB()
	if err != nil {
		fmt.Println(err)
	}

	clubdata := db.GetCollection("clubdata", &client)
	staddata := db.GetCollection("stadiums", &client)

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

	//handling /clubs/get/all
	stadA := Stadiums.StadiumAll{
		Collection: staddata,
	}
	go http.Handle("/stad/get/all", &stadA)

	//handling /stad/get/all/{limit}
	stadAL := Stadiums.StadiumAllLimit{
		Collection: staddata,
	}
	go http.Handle("/stad/get/all/{limit}", &stadAL)

	http.ListenAndServe(":8080", nil)
}
