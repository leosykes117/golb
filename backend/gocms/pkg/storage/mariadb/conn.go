package mariadb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var once sync.Once

// NewConn singleton para crear el pool de conecciones a la base de datos MariaDB
func NewConn() (db *sql.DB, err error) {
	once.Do(func() {
		dbHost := os.Getenv("GOAPI_DB_HOST")
		dbPort := os.Getenv("GOAPI_DB_PORT")
		dbUser := os.Getenv("GOAPI_DB_USER")
		dbPass := os.Getenv("GOAPI_DB_PASSWORD")
		dbName := os.Getenv("GOAPI_DB_NAME")

		if dbHost == "" {
			dbHost = "localhost"
		}
		if dbPort == "" {
			dbPort = "3307"
		}
		if dbUser == "" {
			dbUser = "golb_dev"
		}
		if dbPass == "" {
			dbPass = "100PerApi"
		}
		if dbName == "" {
			dbName = "golb"
		}

		connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
		db, err = sql.Open("mysql", connStr)
		if err != nil {
			log.Fatalf("No se puede conectar la base de datos: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("No hay comunicación con la base de datos: %v", err)
		}

		log.Println("Conectado a MariaDB")
	})
	return
}
