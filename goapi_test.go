package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"
)

var baseUrl string = "http://localhost:8080/v1/"

func TestGetUser(t *testing.T) {
	resp, err := http.Get(baseUrl + "users")

	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		t.Error("Non 200 status code was received!")
	}
}

func TestPostUser(t *testing.T) {
	var jsonStr = []byte(`{"username":"test", "password":"testpass", "firstname":"Tester", "lastname":"Testovich"}`)

	req, _ := http.NewRequest("POST", baseUrl+"users", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Unable to POST new user!")
	}
}

func TestGetUserDetail(t *testing.T) {
	resp, err := http.Get(baseUrl + "users/" + "1")
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		fmt.Println("Unable to fetch specific user by id!")
	}
}
