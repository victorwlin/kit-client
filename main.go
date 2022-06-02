package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set")
	}

	r := mux.NewRouter()

	r.Handle("/favicon.ico", http.NotFoundHandler())

	r.HandleFunc("/", index)
	r.HandleFunc("/friend/", showFriend)
	r.HandleFunc("/addfriend/", addFriend)
	r.HandleFunc("/editfriend/", editFriend)
	r.HandleFunc("/deletefriend/", deleteFriend)

	http.Handle("/", r)

	http.ListenAndServe(":"+port, nil)
}
