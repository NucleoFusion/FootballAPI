package auth

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"api.com/example/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Collection *mongo.Collection
}

func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	queries := r.URL.Query()

	email, name := extractData(queries)
	key := newKey()

	user := models.UserData{
		Email: email,
		Name:  name,
		Key:   key,
	}

	err := insertIntoUsers(a.Collection, &user)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	data, _ := json.Marshal(user)

	io.Writer.Write(w, data)
}

func extractData(queries url.Values) (string, string) {
	return queries["email"][0], queries["name"][0]
}

func newKey() string {
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 5)

	if _, err := src.Read(b); err != nil {
		panic(err)
	}

	hashed, _ := bcrypt.GenerateFromPassword(b, 5)

	return string(hashed)
}

func insertIntoUsers(coll *mongo.Collection, user *models.UserData) error {
	userData := bson.M{
		"email": user.Email,
		"name":  user.Name,
		"key":   user.Key,
	}

	userFound := models.UserData{}
	ExistsErr := coll.FindOne(context.Background(), bson.M{
		"email": user.Email,
	}).Decode(&userFound)
	if ExistsErr != mongo.ErrNoDocuments {
		return errors.New("Already Exists")
	}

	_, err := coll.InsertOne(context.Background(), userData)
	if err != nil {
		return err
	}

	return nil
}
