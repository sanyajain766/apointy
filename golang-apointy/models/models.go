package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type User struct {
		ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		Name   string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,name"`
		Email  string             `json:"email" bson:"email,omitempty" validate:"required,email"`
		Password string           `json:"password" bson:"password,omitempty" validate:"required,password"`		  
	   }

type Post struct {
	ID 			    primitive.ObjectID    `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption    	       string	 			`json:"caption,omitempty" bson:"caption,omitempty"`
	ImageURL  		   string  				`json:"imageurl,omitempty" bson:"imageurl,omitempty" validate:"required,imageurl"`
	PostedTimestamp    time.Time 				`json:"postedtimestamp,omitempty" bson:"postedtimestamp,omitempty" validate:"required,postedtimestamp" validate:"datetime"`
	UserID				string				
}
