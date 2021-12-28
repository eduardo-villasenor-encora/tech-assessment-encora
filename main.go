package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"
	"tech-assessment/clients"
	"tech-assessment/services"
)

func main() {
	http.HandleFunc("/v1/user-posts/", GetUserPostsHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

var GetUserPostsHandler = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("we only allow get method")); err != nil {
			// Logrus to be implemented for better logging
			log.Println("ERROR: we cannot write on response body")
		}
		return
	}

	fullPath := r.URL.Path
	pathUserID := path.Base(fullPath)
	userID, err := strconv.Atoi(pathUserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userService := services.UserService{
		UserClient: clients.UsersClient{
			Client:     http.DefaultClient,
			BaseDomain: "https://jsonplaceholder.typicode.com",
		},
		PostClient: clients.PostsClient{
			Client:     http.DefaultClient,
			BaseDomain: "https://jsonplaceholder.typicode.com",
		},
	}
	userServiceResponse, err := userService.GetUserPosts(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	responseBody, err := json.Marshal(userServiceResponse)
	if err != nil {
		log.Printf("cannot marshal data for response %v due : %s", responseBody, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseBody)
	if err != nil {
		log.Printf("cannot write the response body %v", responseBody)
		return
	}
	return
}
