package models

import (
	"encoding/json"
	"io"
)

type Login struct {
	email    string
	password string
}

func NewLogIn(e, pass string) *Login {
	return &Login{
		email:    e,
		password: pass,
	}
}

func (l *Login) GetEmail() string {
	return l.email
}

func (l *Login) SetEmail(e string) {
	l.email = e
}

func (l *Login) GetPassword() string {
	return l.password
}

func (l *Login) SetPassword(p string) {
	l.password = p
}

func (l *Login) Marshal() ([]byte, error) {
	tmpSrct := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    l.email,
		Password: l.password,
	}

	return json.MarshalIndent(tmpSrct, "", "\t")
}

func (l *Login) Unmarshal(r io.Reader) error {
	tmpSrct := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(r).Decode(&tmpSrct)
	if err != nil {
		return err
	}
	l.email = tmpSrct.Email
	l.password = tmpSrct.Password
	return nil
}
