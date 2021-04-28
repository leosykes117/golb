package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leosykes117/golb/backend/gocms/pkg/models/user"
	"github.com/leosykes117/golb/backend/gocms/pkg/storage/mariadb"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

type Driver string

const (
	MariaDB Driver = "MARIADB"
	MySQL   Driver = "MYSQL"
)

// New crea la conexión con la base de datos
func New(d Driver) {
	switch d {
	case MariaDB:
		db, _ = mariadb.NewConn()
	default:
		log.Fatalf("El driver %s no esta implementado", d)
	}
}

// Pool retorna una unica instancia de DB
func Pool() *sql.DB {
	return db
}

// DAOUser es la fabrica de user.Repository
func DAOUser(driver Driver) (user.Repository, error) {
	switch driver {
	case MariaDB:
		return mariadb.NewRepository(db), nil
	default:
		return nil, fmt.Errorf("El driver %v no está implementado actualmente", driver)
	}
}
