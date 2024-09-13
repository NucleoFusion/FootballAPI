package Stadiums

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"api.com/example/models"
	"api.com/example/routes/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StadiumQuery struct {
	Collection *mongo.Collection
	UserData   *mongo.Collection
}

func (c *StadiumQuery) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	res, err := findQueried(c.Collection, mappedQueries)
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

func handleQueries(queries url.Values) (bson.M, error) {
	m := bson.M{}
	for key, value := range queries {
		if key == "Conf" {
			m["Confederation"] = value[0]
		} else if key == "Country" {
			m["IOC"] = value[0]
		}
	}
	fmt.Println(m)
	return m, nil
}

func findQueried(coll *mongo.Collection, m bson.M) (*mongo.Cursor, error) {
	res, err := coll.Find(context.Background(), m)
	if err != nil {
		return res, err
	}

	return res, nil
}
