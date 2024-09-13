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
)

type StadiumSortLimit struct {
	Collection *mongo.Collection
	UserData   *mongo.Collection
}

func (c *StadiumSortLimit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	sortBy := r.PathValue("sortVal")
	asc := r.URL.Query().Get("asc")

	res, err := findAll(c.Collection)
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
