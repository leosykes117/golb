package main

import (
	"log"

	"github.com/leosykes117/golb/backend/gocms/pkg/models/user"
	"github.com/leosykes117/golb/backend/gocms/pkg/storage"
)

func main() {
	dbDriver := storage.MariaDB
	storage.New(dbDriver)
	repo, err := storage.DAOUser(dbDriver)
	if err != nil {
		log.Fatalf("DAOUser: %v", err)
	}

	m := user.New("leo.aremtz98@gmail.com", "maleva123", "Leonardo Antonio", "Arellano Martínez", "M", "5514670431")

	if err := repo.Create(m); err != nil {
		log.Fatalf("ERROR MariaDB on user.Create: %v", err)
	}
}
