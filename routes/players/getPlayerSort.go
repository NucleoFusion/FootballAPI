package players

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"api.com/example/models"
	"api.com/example/routes/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayerSort struct {
	Collection *mongo.Collection
	UserData   *mongo.Collection
}

func (c *PlayerSort) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := r.URL.Query().Get("key")
	_, err := auth.AuthenticateKey(key, c.UserData)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	sortBy := r.PathValue("sortVal")
	asc := r.URL.Query().Get("asc")
	var intAsc int64
	if asc == "false" {
		intAsc = -1
	} else {
		intAsc = 1
	}

	m := bson.M{}
	m[sortBy] = intAsc

	fmt.Println(m)

	opts := options.Find().SetSort(m)

	res, err := findAll(c.Collection, opts)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	arr := []models.PlayerData{}

	for res.Next(context.Background()) {
		r := models.PlayerData{}

		err := res.Decode(&r)
		if err != nil {
			io.WriteString(w, err.Error())
		}

		arr = append(arr, r)
	}

	data, _ := json.Marshal(arr)

	io.Writer.Write(w, data)
}
