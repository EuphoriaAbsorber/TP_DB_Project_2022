package usecase

import (
	"dbproject/model"
	rep "dbproject/repository"
	"time"
)

type UsecaseInterface interface {
	CreateUser(params *model.User) error
	GetUsersByUsermodel(params *model.User) ([]*model.User, error)
	GetProfile(nickname string) (*model.User, error)
	ChangeProfile(params *model.User) error
	CreateForum(params *model.Forum) error
	GetForumByUsername(nickname string) (*model.Forum, error)
	GetForumBySlug(slug string) (*model.Forum, error)
	GetThreadByModel(params *model.Thread) (*model.Thread, error)
	CreateThreadByModel(params *model.Thread) (*model.Thread, error)
	GetForumUsers(slug string, limit int, since string, desc bool) ([]*model.User, error)
	GetForumThreads(slug string, limit int, since time.Time, desc bool) ([]*model.Thread, error)
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
func (api *Usecase) CreateForum(params *model.Forum) error {
	return api.store.CreateForum(params)
}
func (api *Usecase) GetForumByUsername(nickname string) (*model.Forum, error) {
	return api.store.GetForumByUsername(nickname)
}
func (api *Usecase) GetForumBySlug(slug string) (*model.Forum, error) {
	return api.store.GetForumBySlug(slug)
}
func (api *Usecase) GetThreadByModel(params *model.Thread) (*model.Thread, error) {
	return api.store.GetThreadByModel(params)
}
func (api *Usecase) CreateThreadByModel(params *model.Thread) (*model.Thread, error) {
	return api.store.CreateThreadByModel(params)
}
func (api *Usecase) GetForumUsers(slug string, limit int, since string, desc bool) ([]*model.User, error) {
	return api.store.GetForumUsers(slug, limit, since, desc)
}
func (api *Usecase) GetForumThreads(slug string, limit int, since time.Time, desc bool) ([]*model.Thread, error) {
	return api.store.GetForumThreads(slug, limit, since, desc)
}
