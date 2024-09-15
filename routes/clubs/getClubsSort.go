package Clubs

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"

	"api.com/example/models"
	"api.com/example/routes/auth"
	"api.com/example/statics"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClubAllSort struct {
	Collection *mongo.Collection
	UserData   *mongo.Collection
}

func (c *ClubAllSort) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	opts := options.Find()

	key := r.URL.Query().Get("key")
	_, err := auth.AuthenticateKey(key, c.UserData)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	r.URL.Query().Del("key")

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
	data, _ := json.Marshal(arr)

	io.Writer.Write(w, data)
}

func sortData(arr []models.ClubData, sortBy string) ([]models.ClubData, error) {
	sortKey := statics.SortVals[sortBy]
	if sortKey == "" {
		return nil, errors.New("invalid sortby value")
	}
	// fmt.Println(reflect.Indirect(reflect.ValueOf(arr[0])).FieldByName(sortKey))
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if sortKey == "Goals" {
				if reflect.Indirect(reflect.ValueOf(arr[j])).FieldByName(sortKey).Interface().(int32) > reflect.Indirect(reflect.ValueOf(arr[j+1])).FieldByName(sortKey).Interface().(int32) {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			} else {
				if reflect.Indirect(reflect.ValueOf(arr[j])).FieldByName(sortKey).Interface().(float64) > reflect.Indirect(reflect.ValueOf(arr[j+1])).FieldByName(sortKey).Interface().(float64) {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}

		}
	}
	return arr, nil
}
