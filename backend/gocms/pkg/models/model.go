package models

import "errors"

var (
	// ErrUserCanNotBeNil error para especificar el el usuario no puede ser nil
	ErrUserCanNotBeNil = errors.New("No se especificarón los datos del usuario")
)
