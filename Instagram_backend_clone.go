package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type user struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username      string    `json:"username" gorm:"unique" bson:"username,omitempty"`
    Email         string    `json:"email" gorm:"unique" bson:"email,omitempty"`
    Password      []byte    `json:"password" bson:"password"`
    CreatedAt     time.Time `json:"createdat" bson:"createat"`
    DeactivatedAt time.Time `json:"updatedat" bson:"updatedat"`
	PostID 	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}

type post struct{
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PostID 	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Postdata         string    `json:"email" gorm:"unique" bson:"email,omitempty"`
}

func CreateuserEndpoint(response http.ResponseWriter, request *http.Request) {}
func GetUsersEndpoint(response http.ResponseWriter, request *http.Request) { }
func GetuserEndpoint(response http.ResponseWriter, request *http.Request) { }
func CreatepostEndpoint(response http.ResponseWriter, request *http.Request){}
func GetpostsEndpoint(response http.ResponseWriter, request *http.Request){}
func GetpostEndpoint(response http.ResponseWriter, request *http.Request){}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateuserEndpoint).Methods("POST")
	router.HandleFunc("/Users", GetUsersEndpoint).Methods("GET")
	router.HandleFunc("/user/{id}", GetuserEndpoint).Methods("GET")
	router.HandleFunc("/post", CreatepostEndpoint).Methods("POST")
	router.HandleFunc("/posts", GetpostsEndpoint).Methods("GET")
	router.HandleFunc("/post/{id}", GetuserEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}

func CreateuserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user user
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("Samrat").Collection("Users") //Enterring users
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
}

func CreatepostEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var post post
	_ = json.NewDecoder(request.Body).Decode(&post)
	collection := client.Database("Samrat").Collection("posts") //Enterring posts
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, post)
	json.NewEncoder(response).Encode(result)
}


func GetuserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user user
	collection := client.Database("Samrat").Collection("Users") //Enter the name of user to be found in the list
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, user{ID: id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(user)
}
func GetpostEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var post post
	collection := client.Database("Samrat").Collection("posts") //Enter the name of user to be found in the list
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Post{PostID: id}).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(post)
}

func GetUsersEndpoint(response http.ResponseWriter, request *http.Request) { //All the users using id
	response.Header().Set("content-type", "application/json")
	var Users []user
	collection := client.Database("Samrat").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user user
		cursor.Decode(&user)
		Users = append(Users, user)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(Users)
}

func GetpostsEndpoint(response http.ResponseWriter, request *http.Request) { //All the posts using id
	response.Header().Set("content-type", "application/json")
	var posts []post
	collection := client.Database("Samrat").Collection("Posts")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post post
		cursor.Decode(&post)
		post = append(posts, post)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(posts)
}
