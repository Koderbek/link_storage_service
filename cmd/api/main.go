package main

import (
	"fmt"
	"github.com/Koderbek/link_storage_service/internal/config"
	"github.com/Koderbek/link_storage_service/internal/database"
	"github.com/Koderbek/link_storage_service/internal/server"
	"net/http"
)

func main() {
	cfg := config.Init()
	db, err := database.NewPostgresDb(cfg)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	repo := database.NewRepository(db)
	srv := server.NewServer(repo)
	if err = http.ListenAndServe(":8080", srv); err != nil {
		fmt.Println(err)
	}
}
