package main

import (
	"fmt"
	"github.com/Koderbek/link_storage_service/internal/config"
	"github.com/Koderbek/link_storage_service/internal/database"
)

func main() {
	cfg := config.Init()
	db, err := database.NewPostgresDb(cfg)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	repo := database.NewRepository(db)

	//TODO: test
	repo.Create("https://github.com/jmoiron/sqlx", "abc123")
	link, err := repo.Link("abc1234")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(link)
}
