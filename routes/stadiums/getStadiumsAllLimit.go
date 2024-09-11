package Stadiums

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"api.com/example/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type StadiumAllLimit struct {
	Collection *mongo.Collection
}

func (c *StadiumAllLimit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	lim := r.PathValue("limit")
	limit, _ := strconv.Atoi(lim)

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

	data, _ := json.Marshal(arr[:limit])

	io.Writer.Write(w, data)
}
