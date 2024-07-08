package database

import "github.com/willbrr.dev/goexpert/9-API/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
