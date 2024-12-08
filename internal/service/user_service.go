package service

import (
	"IMChat_App/internal/model"
	"IMChat_App/internal/repository"
	"IMChat_App/pkg/common"
)

func Register(user *model.User) bool {
	return repository.Register(user)
}

func Login(user *model.User) bool {
	return repository.Login(user)
}

func SearchUser(account string) (bool, *model.User) {
	return repository.SearchUser(account)
}

func AddUser(relation *common.AddUserReq) {
	repository.AddUser(relation)
}
