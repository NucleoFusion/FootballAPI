package Stadiums

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"api.com/example/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StadiumAll struct {
	Collection *mongo.Collection
}

func (c *StadiumAll) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
