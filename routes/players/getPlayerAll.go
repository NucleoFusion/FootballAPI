package players

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"api.com/example/models"
	"api.com/example/routes/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayersAll struct {
	Collection *mongo.Collection
	UserData   *mongo.Collection
}

func (c *PlayersAll) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	opts := options.Find()

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

func findAll(coll *mongo.Collection, opts *options.FindOptions) (*mongo.Cursor, error) {
	res, err := coll.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return res, err
	}

	return res, nil
}
