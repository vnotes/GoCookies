package main

import (
	"net/http"

	"github.com/vnotes/gocookies/service/user/api"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/user", api.GetUserInfo).Methods(http.MethodGet)
	r.HandleFunc("/api/user", api.CreateUser).Methods(http.MethodPost)

	_ = http.ListenAndServe(":11111", r)
}
