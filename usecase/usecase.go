package usecase

import (
	"dbproject/model"
	rep "dbproject/repository"
)

type UsecaseInterface interface {
	CreateUser(params *model.User) error
	GetUsersByUsermodel(params *model.User) ([]*model.User, error)
	GetProfile(nickname string) (*model.User, error)
	ChangeProfile(params *model.User) error
}

type Usecase struct {
	store rep.StoreInterface
}

func NewUsecase(s rep.StoreInterface) UsecaseInterface {
	return &Usecase{
		store: s,
	}
}

func (api *Usecase) CreateUser(params *model.User) error {
	return api.store.CreateUser(params)
}
func (api *Usecase) GetUsersByUsermodel(params *model.User) ([]*model.User, error) {
	return api.store.GetUsersByUsermodel(params)
}
func (api *Usecase) GetProfile(nickname string) (*model.User, error) {
	return api.store.GetProfile(nickname)
}
func (api *Usecase) ChangeProfile(params *model.User) error {
	return api.store.ChangeProfile(params)
}
