package main

import (
	"net/http"

	"github.com/ShinnosukeSuzuki/web-app-develop-golang-todo/handler"
	"github.com/ShinnosukeSuzuki/web-app-develop-golang-todo/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux() http.Handler {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	v := validator.New()
	at := &handler.AddTask{Store: store.Tasks, Validator: v}
	mux.Post("/tasks", at.ServeHttp)
	lt := &handler.ListTask{Store: store.Tasks}
	mux.Get("/tasks", lt.ServeHttp)
	return mux
}
