package api

import (
	"log"

	"github.com/leosykes117/golb/backend/gocms/pkg/models/user"
	"github.com/leosykes117/golb/backend/gocms/pkg/storage"
)

type Services struct {
	userService *user.Service
}

func Start(port string) {
	log.Println("Start API PORT:", port)

	driver := storage.MariaDB
	storage.New(driver)

	mariaStorageProduct, err := storage.DAOUser(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}

	services := &Services{
		userService: user.NewService(mariaStorageProduct),
	}

	r := routes(services)
	server := newServer(port, r)
	server.Start()
}
