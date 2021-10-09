package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	helper "github.com/sanyajain756/golang-appointy/helper"
)
var Usercollection = helper.ConnectDB1()
var Postcollection = helper.ConnectDB2()
func TestGetUser(t *testing.T) {
		req, err := http.NewRequest("GET", "/users/61615af60647ae7f1937bc52", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(getUser)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	
		
	
	}



	func TestGetPost(t *testing.T) {
		req, err := http.NewRequest("GET", "/posts/61617a153ac91a52f1120547", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(getPost)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	
		
	
	}

func TestCreateUser(t *testing.T) {

	var jsonStr = []byte(`{"_id":"61615af60647ae7f1937bc52","name":"sanya","email":"21ws","password":"$2a$14$f6ndzUtt9ZelDpljIM2jue23xfXqIcc7a8E1bM86Nrx2PAa6T23Nm"}`)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
func TestCreatePost(t *testing.T) {

	var jsonStr = []byte(`{"_id":"61617a153ac91a52f1120547","caption":"hey there","imageurl":"12123e32","postedtimestamp":"2021-10-09T11:16:37.709+00:00","userid":"edewed"}`)

	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	
}
func TestGetPosts(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/users/edewed", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getPosts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	
	

}
