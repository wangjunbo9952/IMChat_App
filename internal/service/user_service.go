package service

import (
	"IMChat_App/internal/model"
	"IMChat_App/internal/repository"
)

func Register(user *model.User) bool {
	return repository.Register(user)
}
