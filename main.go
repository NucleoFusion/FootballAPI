package main

import (
	"fmt"
	"net/http"

	"api.com/example/db"
	"api.com/example/routes/auth"
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
	userData := db.GetCollection("users", &client)

	//AUTH

	auth := auth.Auth{
		Collection: userData,
	}
	go http.Handle("/register", &auth)

	//CLUBS

	// handling /clubs/get
	c := Clubs.ClubHandler{
		Collection: clubdata,
		UserData:   userData,
	}
	go http.Handle("/clubs/get/all", &c)

	//handling /clubs/get/query
	cq := Clubs.ClubQuery{
		Collection: clubdata,
		UserData:   userData,
	}
	go http.Handle("/clubs/get/query", &cq)

	//handling /clubs/get/all/limit
	cal := Clubs.ClubAllLimit{
		Collection: clubdata,
		UserData:   userData,
	}
	go http.Handle("/clubs/get/all/{limit}", &cal)

	//handling /clubs/get/query/limit
	cql := Clubs.ClubAllLimit{
		Collection: clubdata,
		UserData:   userData,
	}
	go http.Handle("/clubs/get/query/{limit}", &cql)

	//handling /clubs/get/sortBy/{val}
	cas := Clubs.ClubAllSort{
		Collection: clubdata,
		UserData:   userData,
	}
	go http.Handle("/clubs/get/sortBy/{sortVal}", &cas)

	//handling /clubs/get/sortBy/{val}/{limit}
	cls := Clubs.ClubLimitSort{
		Collection: clubdata,
		UserData:   userData,
	}
	go http.Handle("/clubs/get/sortBy/{sortVal}/limit/{limit}", &cls)

	//handling /clubs/get/all
	stadA := Stadiums.StadiumAll{
		Collection: staddata,
		UserData:   userData,
	}
	go http.Handle("/stad/get/all", &stadA)

	//handling /stad/get/all/{limit}
	stadAL := Stadiums.StadiumAllLimit{
		Collection: staddata,
		UserData:   userData,
	}
	go http.Handle("/stad/get/all/{limit}", &stadAL)

	//handling /stad/get/query
	stadQ := Stadiums.StadiumQuery{
		Collection: staddata,
		UserData:   userData,
	}
	go http.Handle("/stad/get/query", &stadQ)

	//handling /stad/get/query/{limit}
	stadQL := Stadiums.StadiumQueryLimit{
		Collection: staddata,
		UserData:   userData,
	}
	go http.Handle("/stad/get/query/{limit}", &stadQL)

	//handling /stad/get/sortBy/{sortVal}
	stadS := Stadiums.StadiumAllSort{
		Collection: staddata,
		UserData:   userData,
	}
	go http.Handle("/stad/get/sortBy/{sortVal}", &stadS)

	//handling /stad/get/sortBy/{sortVal}/{limit}
	stadSL := Stadiums.StadiumSortLimit{
		Collection: staddata,
		UserData:   userData,
	}
	go http.Handle("/stad/get/sortBy/{sortVal}/{limit}", &stadSL)

	http.ListenAndServe(":8080", nil)
}
