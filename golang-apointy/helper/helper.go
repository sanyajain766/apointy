package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB1() *mongo.Collection {

	clientOptions := options.Client().ApplyURI("mongodb+srv://sanyajain:sanyajain@instagram.58jbi.mongodb.net/instagram?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("instagram").Collection("users")
	

	return collection
}
func ConnectDB2() *mongo.Collection {


	clientOptions := options.Client().ApplyURI("mongodb+srv://sanyajain:sanyajain@instagram.58jbi.mongodb.net/instagram?retryWrites=true&w=majority")


	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection1 := client.Database("instagram").Collection("posts")


	return collection1
}



type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}