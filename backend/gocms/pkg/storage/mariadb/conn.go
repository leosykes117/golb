package mariadb

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leosykes117/golb/backend/gocms/internal/env"
	_ "github.com/lib/pq"
)

var once sync.Once

// NewConn singleton para crear el pool de conecciones a la base de datos MariaDB
func NewConn() (db *sql.DB, err error) {
	once.Do(func() {
		dbHost, _ := env.GetEnvs(env.DBHost)
		dbPort, _ := env.GetEnvs(env.DBPort)
		dbUser, _ := env.GetEnvs(env.DBUser)
		dbPass, _ := env.GetEnvs(env.DBPassword)
		dbName, _ := env.GetEnvs(env.DBName)

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
