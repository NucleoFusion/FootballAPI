package Clubs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"api.com/example/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClubQueryLimit struct {
	Collection *mongo.Collection
}

func (c *ClubQueryLimit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	lim := r.PathValue("limit")
	limit, _ := strconv.Atoi(lim)

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

	data, _ := json.Marshal(arr[:limit])

	io.Writer.Write(w, data)
}
