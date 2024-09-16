package players

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"api.com/example/models"
	"api.com/example/routes/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayerSortLimit struct {
	Collection *mongo.Collection
	UserData   *mongo.Collection
}

func (c *PlayerSortLimit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	lim := r.PathValue("limit")
	limit, _ := strconv.Atoi(lim)

	m := bson.M{}
	m[sortBy] = intAsc

	opts := options.Find().SetSort(m).SetLimit(int64(limit))

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
