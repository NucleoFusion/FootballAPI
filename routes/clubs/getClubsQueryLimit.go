package Clubs

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"api.com/example/models"
	"api.com/example/routes/auth"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClubQueryLimit struct {
	Collection *mongo.Collection
	UserData   *mongo.Collection
}

func (c *ClubQueryLimit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	lim := r.PathValue("limit")
	limit, _ := strconv.Atoi(lim)

	key := r.URL.Query().Get("key")
	_, err := auth.AuthenticateKey(key, c.UserData)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	r.URL.Query().Del("key")

	queries := r.URL.Query()

	mappedQueries, err := handleQueries(queries)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	res, err := findQueried(c.Collection, mappedQueries)
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

	data, _ := json.Marshal(arr[:limit])

	io.Writer.Write(w, data)
}
