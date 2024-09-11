package Clubs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"api.com/example/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClubHandler struct {
	Collection *mongo.Collection
}

func (c *ClubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

// func readCursorClub(res *mongo.Cursor) (*[]models.ClubData, error) {
// 	arr := make([]models.ClubData, 98)

// 	for res.Next(context.Background()) {
// 		r := models.ClubData{}

// 		err := res.Decode(&r)
// 		fmt.Println(r)
// 		if err != nil {
// 			return &arr, err
// 		}

// 		arr = append(arr, r)
// 	}

// 	return &arr, nil
// }
