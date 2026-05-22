package main

import (
	"fmt"
	"log"
	"net/http"

	"your-module-name/internal/config"
	"your-module-name/internal/database"
)

func main() {
	cfg := config.Load()

	db := database.NewPostgres(cfg)
	defer db.Close()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	log.Printf("Server running on :%s", cfg.Port)

	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}