package typesconv

import "github.com/leosykes117/golb/backend/gocms/pkg/models/user"

type scanner interface {
	Scan(dest ...interface{}) error
}

func ScanRowUser(s scanner) (*user.Model, error) {
	var (
		id      uint
		email   string
		name    string
		surname string
		gender  string
		phone   string
	)

	err := s.Scan(&id, &email, &name, &surname, &gender, &phone)
	if err != nil {
		return nil, err
	}

	m := user.New(email, "", name, surname, gender, phone)
	m.ID = id

	return m, nil
}
