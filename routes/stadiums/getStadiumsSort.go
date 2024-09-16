package Stadiums

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"api.com/example/models"
	"api.com/example/routes/auth"
	"api.com/example/statics"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StadiumAllSort struct {
	Collection *mongo.Collection
	UserData   *mongo.Collection
}

func (c *StadiumAllSort) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := r.URL.Query().Get("key")
	_, err := auth.AuthenticateKey(key, c.UserData)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	sortBy := r.PathValue("sortVal")
	asc := r.URL.Query().Get("asc")
	var intAsc int64
	if asc == "false" {
		intAsc = -1
	} else {
		intAsc = 1
	}

	m := bson.M{}
	m[sortBy] = intAsc

	opts := options.Find().SetSort(m)

	res, err := findAll(c.Collection, opts)
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
	data, _ := json.Marshal(arr)

	io.Writer.Write(w, data)
}

func sortData(arr []models.StadiumData, sortBy string) ([]models.StadiumData, error) {
	sortKey := statics.SortValsStad[sortBy]
	fmt.Println(sortKey)
	if sortKey == "" {
		return nil, errors.New("invalid sortby value")
	}
	// fmt.Println(reflect.Indirect(reflect.ValueOf(arr[0])).FieldByName(sortKey))
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if reflect.Indirect(reflect.ValueOf(arr[j])).FieldByName(sortKey).Interface().(int32) > reflect.Indirect(reflect.ValueOf(arr[j+1])).FieldByName(sortKey).Interface().(int32) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr, nil
}
