package user

import (
	"fmt"

	"github.com/leosykes117/golb/backend/gocms/pkg/models"
)

// Model es la estructura que representa un usuario.
type Model struct {
	ID           uint   `json:"ID"`
	Email        string `json:"email,omitempty"`
	PasswordHash string `json:"password,omitempty"`
	Name         string `json:"name,omitempty"`
	Surname      string `json:"surname,omitempty"`
	Gender       string `json:"gender,omitempty"`
	Phone        string `json:"phone,omitempty"`
}

// Models es un array de punteros de model
type Models []*Model

// New retorna un puntero a una nueva instancia de Model
func New(email, password, name, surname, gender, phone string) *Model {
	return &Model{
		Email:        email,
		PasswordHash: password,
		Name:         name,
		Surname:      surname,
		Gender:       gender,
		Phone:        phone,
	}
}

func (u *Model) FullName() string {
	return fmt.Sprintf("%s %s", u.Name, u.Surname)
}

// Repository es la interfaz que implementa los metodos para la base de datos
type Repository interface {
	Create(*Model) error
	LogIn(string, string) (*Model, error)
	//Update(*Model) error
	//GetByID(uint) (*Model, error)
	//Delete(uint) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(m *Model) error {
	return s.repo.Create(m)
}

func (s *Service) LogIn(lg *models.Login) (*Model, error) {
	return s.repo.LogIn(lg.GetEmail(), lg.GetPassword())
}
