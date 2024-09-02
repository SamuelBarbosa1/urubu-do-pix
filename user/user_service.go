package user

import (
	"errors"
	"fmt"
)

func CreateUser(name, password string) (*User, error) {
	// Simula a criação de uma chave Pix única
	pixKey := fmt.Sprintf("pix-%s", name)
	if _, exists := GetUserByPixKey(pixKey); exists {
		return nil, errors.New("usuário já existe")
	}

	user := &User{
		ID:       fmt.Sprintf("%d", len(users)+1),
		Name:     name,
		Password: password,
		PixKey:   pixKey,
		Balance:  0,
		History:  []string{},
	}
	users[pixKey] = user
	return user, nil
}

func LoginUser(pixKey, password string) (*User, error) {
	user, exists := GetUserByPixKey(pixKey)
	if !exists {
		return nil, errors.New("usuário não encontrado")
	}

	if user.Password != password {
		return nil, errors.New("senha incorreta")
	}

	return user, nil
}

func (u *User) UpdateBalance(amount float64, historyMessage string) {
	u.Balance += amount
	u.History = append(u.History, historyMessage)
	UpdateUser(u)
}
