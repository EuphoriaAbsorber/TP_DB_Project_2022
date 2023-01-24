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
	GetPostById(id int) (*model.Post, error)
	UpdatePost(id int, mes string) (*model.Post, error)
	GetServiceStatus() (*model.Status, error)
	ServiceClear() error
	GetThread(id int, slug string) (*model.Thread, error)
	CreatePosts(in *model.Posts, id int, slug string) ([]*model.Post, error)
	UpdateThreadInfo(in *model.ThreadUpdate, id int, slug string) (*model.Thread, error)
	VoteForThread(in *model.Vote, id int, slug string) (*model.Thread, error)
	GetThreadPosts(slug string, id int, limit int, since int, sort string, desc bool) ([]*model.Post, error)
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
func (api *Usecase) GetPostById(id int) (*model.Post, error) {
	return api.store.GetPostById(id)
}
func (api *Usecase) UpdatePost(id int, mes string) (*model.Post, error) {
	return api.store.UpdatePost(id, mes)
}
func (api *Usecase) GetServiceStatus() (*model.Status, error) {
	return api.store.GetServiceStatus()
}
func (api *Usecase) ServiceClear() error {
	return api.store.ServiceClear()
}

func (api *Usecase) GetThread(id int, slug string) (*model.Thread, error) {
	var thread *model.Thread
	var err error
	if id == 0 {
		thread, err = api.store.GetThreadBySlug(slug)
		if err != nil {
			return nil, err
		}
	} else {
		thread, err = api.store.GetThreadById(id)
		if err != nil {
			return nil, err
		}
	}
	return thread, nil
}

func (api *Usecase) CreatePosts(in *model.Posts, id int, slug string) ([]*model.Post, error) {
	thread, err := api.GetThread(id, slug)
	if err != nil {
		return nil, err
	}
	parentPostsNumbers := make([]int, 0)
	m := make(map[int]bool)
	for _, post := range *in {
		if _, ok := m[post.Parent]; !ok {
			m[post.Parent] = true
			if post.Parent != 0 {
				parentPostsNumbers = append(parentPostsNumbers, post.Parent)
			}
		}
		user, err := api.store.GetProfile(post.Author)
		if err != nil {
			return nil, err
		}
		post.Author = user.Nickname
	}
	err = api.store.CheckAllPostParentIds(thread.Id, parentPostsNumbers)
	if err != nil {
		return nil, err
	}
	return api.store.CreatePosts(in, thread.Id, thread.Forum)
}

func (api *Usecase) UpdateThreadInfo(in *model.ThreadUpdate, id int, slug string) (*model.Thread, error) {
	thread, err := api.GetThread(id, slug)
	if err != nil {
		return nil, err
	}
	if in.Message == "" {
		in.Message = thread.Message
	}
	if in.Title == "" {
		in.Title = thread.Title
	}
	err = api.store.UpdateThreadInfo(in, thread.Id)
	if err != nil {
		return nil, err
	}
	thread.Message = in.Message
	thread.Title = in.Title
	return thread, nil
}

func (api *Usecase) VoteForThread(in *model.Vote, id int, slug string) (*model.Thread, error) {
	thread, err := api.GetThread(id, slug)
	if err != nil {
		return nil, err
	}
	user, err := api.store.GetProfile(in.Nickname)
	if err != nil {
		return nil, err
	}
	in.Nickname = user.Nickname
	newRating, err := api.store.VoteForThread(in, thread.Id)
	if err != nil {
		return nil, err
	}
	thread.Votes = newRating
	return thread, nil
}
func (api *Usecase) GetThreadPosts(slug string, id int, limit int, since int, sort string, desc bool) ([]*model.Post, error) {
	thread, err := api.GetThread(id, slug)
	if err != nil {
		return nil, err
	}
	return api.store.GetThreadPosts(thread.Id, limit, since, sort, desc)
}
