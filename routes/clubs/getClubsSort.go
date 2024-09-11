package Clubs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"api.com/example/models"
	"api.com/example/statics"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClubAllSort struct {
	Collection *mongo.Collection
}

func (c *ClubAllSort) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sortBy := r.PathValue("sortVal")
	asc := r.URL.Query().Get("asc")

	res, err := findAll(c.Collection)
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

	arr = sortData(arr, sortBy)
	if asc == "false" {
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	data, _ := json.Marshal(arr)

	io.Writer.Write(w, data)
}

func sortData(arr []models.ClubData, sortBy string) []models.ClubData {
	sortKey := statics.SortVals[sortBy]
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
	return arr
}
