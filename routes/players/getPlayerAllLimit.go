package players

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

type PlayersAllLimit struct {
	Collection *mongo.Collection
	UserData   *mongo.Collection
}

func (c *PlayersAllLimit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	res, err := findAll(c.Collection, opts)
	if err != nil {
		io.WriteString(w, err.Error())
		return
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
