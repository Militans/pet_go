package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	db, err := sql.Open("MariaDB", "root:1q2w_1q2w@tcp(localhost:3306)/for_go")
	if err != nil {
		fmt.Println("Failed connect")
	}
	defer db.Close()
	
	router := chi.NewRouter()

	router.Get("/hello", basicHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
