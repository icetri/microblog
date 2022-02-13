package server

import (
	"github.com/gorilla/mux"
	"microblog/server/handlers"
	"net/http"
)

func NewRouter(router *mux.Router, h *handlers.Handler) {

	router.HandleFunc("/", h.Index).Methods("GET")
	router.HandleFunc("/aboutus/", h.AboutUs).Methods("GET")
	router.HandleFunc("/blog/", h.Blog).Methods("GET")
	router.HandleFunc("/saveparamsforblog/", h.SaveParamsFromBlog).Methods("POST")
	router.HandleFunc("/post/{id:[0-9]+}", h.ShowPost).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
}
