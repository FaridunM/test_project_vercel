package main

import (
	"fmt"
	"net/http"
	handler "test_project_vercel/api"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Обработчик маршрута
	r.HandleFunc("/", handler.Handler).Methods("POST")

	fmt.Println("Сервер запущен на http://0.0.0.0:8080")
	http.ListenAndServe(":8080", r)
}
