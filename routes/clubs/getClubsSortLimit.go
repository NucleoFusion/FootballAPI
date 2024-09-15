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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClubLimitSort struct {
	Collection *mongo.Collection
	UserData   *mongo.Collection
}

func (c *ClubLimitSort) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	opts := options.Find()

	key := r.URL.Query().Get("key")
	_, err := auth.AuthenticateKey(key, c.UserData)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	r.URL.Query().Del("key")

	lim := r.PathValue("limit")
	limit, _ := strconv.Atoi(lim)

	sortBy := r.PathValue("sortVal")
	asc := r.URL.Query().Get("asc")

	res, err := findAll(c.Collection, opts)
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

	arr, err = sortData(arr, sortBy)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	if asc == "false" {
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	data, _ := json.Marshal(arr[:limit])

	io.Writer.Write(w, data)
}
