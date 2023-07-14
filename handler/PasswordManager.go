package handler

import (
	"errors"
)

var ErrNotUsernameFound = errors.New("no username found")

type PasswordManager interface {
	GetPassword(string) (string, error)
}

type UnsecurePasswordManager struct {
	passwords map[string]string
}

func NewUnsecurePasswordManager() *UnsecurePasswordManager {
	return &UnsecurePasswordManager{
		passwords: map[string]string{},
	}
}

func (p *UnsecurePasswordManager) GetPassword(username string) (string, error) {
	pass, ok := p.passwords[username]
	if !ok {
		return "", ErrNotUsernameFound
	}

	return pass, nil
}

func (p *UnsecurePasswordManager) SetPassword(username, password string) {
	p.passwords[username] = password
}
