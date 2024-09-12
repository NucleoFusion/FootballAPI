package Stadiums

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"api.com/example/models"
	"api.com/example/routes/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StadiumAll struct {
	Collection *mongo.Collection
	UserData   *mongo.Collection
}

func (c *StadiumAll) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := r.URL.Query().Get("key")
	_, err := auth.AuthenticateKey(key, c.UserData)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	r.URL.Query().Del("key")

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

	data, _ := json.Marshal(arr)

	io.Writer.Write(w, data)
}

func findAll(coll *mongo.Collection) (*mongo.Cursor, error) {
	res, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		return res, err
	}

	return res, nil
}
