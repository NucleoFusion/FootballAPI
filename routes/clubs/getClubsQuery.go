package Clubs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"api.com/example/models"
	"api.com/example/statics"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClubQuery struct {
	Collection *mongo.Collection
}

func (c *ClubQuery) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queries := r.URL.Query()

	mappedQueries := handleQueries(queries)

	res, err := findQueried(c.Collection, mappedQueries)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("&v", err))
	}

	arr := []models.ClubData{}

	for res.Next(context.Background()) {
		r := models.ClubData{}

		err := res.Decode(&r)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("&v", err))
		}

		arr = append(arr, r)
	}

	data, _ := json.Marshal(arr)

	io.Writer.Write(w, data)
}

func findQueried(coll *mongo.Collection, queries bson.M) (*mongo.Cursor, error) {
	res, err := coll.Find(context.Background(), queries)
	if err != nil {
		return res, err
	}

	return res, nil
}

func handleQueries(queries url.Values) bson.M {
	m := bson.M{}
	var newValue string
	for key, value := range queries {
		if key == "Tournament" {
			newValue = statics.ClubTournaments[value[0]]
		} else if key == "Team" {
			fmt.Println(key)
			newValue = statics.ClubNames[value[0]]
		}

		m[key] = newValue
	}
	fmt.Println(m)
	return m
}
