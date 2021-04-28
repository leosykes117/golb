package mariadb

import (
	"database/sql"
	"log"

	"github.com/leosykes117/golb/backend/gocms/pkg/models/user"
	"github.com/leosykes117/golb/backend/gocms/pkg/storage/typesconv"
)

const (
	mariaDBCreateUser = `INSERT INTO users (email, password_hash, name, surname, gender, phone)
						VALUES(?, ?, ?, ?, ?, ?) 
						RETURNING user_id`
)

// userRepository es la estrucutra encargada de trabajar la tabla product en postgres
type userRepository struct {
	db *sql.DB
}

// NewRepository crear una instancia de mySQLProduct
func NewRepository(db *sql.DB) user.Repository {
	return &userRepository{db}
}

// Create implementa la interfaz product.Create. Almacena un producto en la db
func (r *userRepository) Create(m *user.Model) error {
	stmt, err := r.db.Prepare(mariaDBCreateUser)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Email,
		m.PasswordHash,
		m.Name,
		m.Surname,
		typesconv.StringToNull(m.Gender),
		typesconv.StringToNull(m.Phone),
	).Scan(&m.ID)
	if err != nil {
		return err
	}

	log.Printf("Se creo el usuario correctamente: %+v", m)
	return nil
}
