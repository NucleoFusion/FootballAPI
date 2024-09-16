package Stadiums

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"api.com/example/models"
	"api.com/example/routes/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StadiumQueryLimit struct {
	Collection *mongo.Collection
	UserData   *mongo.Collection
}

func (c *StadiumQueryLimit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	lim := r.PathValue("limit")
	limit, _ := strconv.Atoi(lim)

	opts := options.Find().SetLimit(int64(limit))

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

	res, err := findQueried(c.Collection, mappedQueries, opts)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	arr := []models.StadiumData{}

	for res.Next(context.Background()) {
		r := models.StadiumData{}

		err := res.Decode(&r)
		if err != nil {
			io.WriteString(w, err.Error())
		}

		arr = append(arr, r)
	}

	data, _ := json.Marshal(arr)

	io.Writer.Write(w, data)
}
