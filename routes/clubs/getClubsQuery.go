package Clubs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"api.com/example/models"
	"api.com/example/routes/auth"
	"api.com/example/statics"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClubQuery struct {
	Collection *mongo.Collection
	UserData   *mongo.Collection
}

func (c *ClubQuery) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := r.URL.Query().Get("key")
	_, err := auth.AuthenticateKey(key, c.UserData)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	queries := r.URL.Query()

	mappedQueries, err := handleQueries(queries)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	opts := options.Find()

	res, err := findQueried(c.Collection, mappedQueries, opts)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	arr := []models.ClubData{}

	for res.Next(context.Background()) {
		r := models.ClubData{}

		err := res.Decode(&r)
		if err != nil {
			io.WriteString(w, err.Error())
		}

		arr = append(arr, r)
	}

	data, _ := json.Marshal(arr)

	io.Writer.Write(w, data)
}

func findQueried(coll *mongo.Collection, queries bson.M, opts *options.FindOptions) (*mongo.Cursor, error) {
	res, err := coll.Find(context.Background(), queries, opts)
	if err != nil {
		return res, err
	}

	return res, nil
}

func handleQueries(queries url.Values) (bson.M, error) {
	m := bson.M{}
	var newValue string
	for key, value := range queries {
		if key == "Tournament" {
			newValue = statics.ClubTournaments[value[0]]
			if newValue == "" {
				return nil, errors.New("invalid Query Params")
			}
		} else if key == "Team" {
			newValue = statics.ClubNames[value[0]]
		} else if key == "key" {
			continue
		}
		m[key] = newValue
	}
	fmt.Println(m)
	return m, nil
}
