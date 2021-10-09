package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	//controllers "github.com/sanyajain756/golang-apointy/controllers"

	helper "github.com/sanyajain756/golang-appointy/helper"
	models "github.com/sanyajain756/golang-appointy/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var (
	Ctx   = context.TODO()
)


var collection = helper.ConnectDB1()
var collection1 = helper.ConnectDB2()
func HashPassword(password string) string {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        log.Panic(err)
    }

    return string(bytes)
}
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method=="POST"{
		var user models.User
		_ = json.NewDecoder(r.Body).Decode(&user)
		if user.Name ==""{
			json.NewEncoder(w).Encode("Name is required")
	
		}
		if user.Email ==""{
			json.NewEncoder(w).Encode("Email is required")
	
		}
		if user.Password ==""{
			json.NewEncoder(w).Encode("Password is required")
	
		}
		
		if user.Name!="" && user.Email!="" && user.Password!=""{
			user.Password=HashPassword(user.Password)
			user:=mongo.IndexModel{
				Keys: bson.M{user.Email: 1}, // index in ascending order or -1 for descending order
				Options: options.Index().SetUnique(true),
			}
			result, err := collection.InsertOne(context.TODO(), user)
			if err != nil {
				helper.GetError(err, w)
				return
			}
			json.NewEncoder(w).Encode(result)
		}
		}else{
			json.NewEncoder(w).Encode("Only POST request is permitted")
	}
	
}
func getUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var user models.User
		params := r.URL.Path[len("/users/"):]
		id, _ := primitive.ObjectIDFromHex(params)
		filter := bson.M{"_id": id}
		err := collection.FindOne(context.TODO(), filter).Decode(&user)
		if err != nil {
			helper.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(user)
	}else{
		json.NewEncoder(w).Encode("Only GET Request is permitted")
	}
}
func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST"{
		var post models.Post
		_ = json.NewDecoder(r.Body).Decode(&post)
		if post.ImageURL ==""{
			json.NewEncoder(w).Encode("ImageURL is required")
			}
		
		if post.UserID ==""{
			json.NewEncoder(w).Encode("UserID is required")
		}
		if post.PostedTimestamp.IsZero() {
			post.PostedTimestamp = time.Now()
		}
		if  post.ImageURL!="" && post.UserID!=""{
			result, err := collection1.InsertOne(context.TODO(), post)
			if err != nil {
				helper.GetError(err, w)
				return
			}
		json.NewEncoder(w).Encode(result)
		}
	}else{
		json.NewEncoder(w).Encode("Only POST Request is permitted")
	}
}
func getPost(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var post models.Post
		var params = r.URL.Path[len("/posts/"):]
		id, _ := primitive.ObjectIDFromHex(params)
		filter := bson.M{"_id": id}
		err := collection1.FindOne(context.TODO(), filter).Decode(&post)
		if err != nil {
			helper.GetError(err, w)
			return
		}
		json.NewEncoder(w).Encode(post)
	}else{
		json.NewEncoder(w).Encode("Only GET request is permitted")
	}
}
func getPosts(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var params = r.URL.Path[len("/posts/users/"):]
		options:=options.Find()
		pages,done:=r.URL.Query()["page"]
		page:=1
		if !done || len(pages[0]) < 1 {
			page=1
		   }else{
			page, _ = strconv.Atoi(pages[0])
		   }
		var perPage int64=5
		options.SetSkip((int64(page)-1)*perPage)
		options.SetLimit(perPage)
		filter,err:=collection1.Find(Ctx,bson.M{"userid":params},options)
		if err != nil {
			helper.GetError(err, w)
			return
		}
		var filtered []bson.M
		if err=filter.All(Ctx,&filtered); err!=nil{
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(filtered)
		}else{
			json.NewEncoder(w).Encode("Only GET request is permitted")
		}
}

func Setup() {
	host := "127.0.0.1"
	port := "27017"
	connectionURI := "mongodb://" + host + ":" + port + "/"
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("instagram")
	collection = db.Collection("users")

}

func main() {
  http.HandleFunc("/users/", getUser)
  http.HandleFunc("/users", createUser)
  http.HandleFunc("/posts", createPost)
  http.HandleFunc("/posts/", getPost)
  http.HandleFunc("/posts/users/", getPosts)
  log.Fatal(http.ListenAndServe(":8000", nil))

}