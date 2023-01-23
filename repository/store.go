package repository

import (
	"context"
	"dbproject/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StoreInterface interface {
	CreateUser(params *model.User) error
	GetUsersByUsermodel(in *model.User) ([]*model.User, error)
	GetProfile(nickname string) (*model.User, error)
	ChangeProfile(in *model.User) error
	CreateForum(in *model.Forum) error
	GetForumByUsername(nickname string) (*model.Forum, error)
	GetForumBySlug(slug string) (*model.Forum, error)
	GetThreadByModel(in *model.Thread) (*model.Thread, error)
	CreateThreadByModel(in *model.Thread) (*model.Thread, error)
	GetForumUsers(slug string, limit int, since string, desc bool) ([]*model.User, error)
}

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) StoreInterface {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateUser(in *model.User) error {
	_, err := s.db.Exec(context.Background(), `INSERT INTO users (email, fullname, nickname, about) VALUES ($1, $2, $3, $4);`, in.Email, in.Fullname, in.Nickname, in.About)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) GetUsersByUsermodel(in *model.User) ([]*model.User, error) {
	users := []*model.User{}
	rows, err := s.db.Query(context.Background(), `SELECT email, fullname, nickname, about FROM users WHERE email = $1 OR nickname = $2;`, in.Email, in.Nickname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		dat := model.User{}
		err := rows.Scan(&dat.Email, &dat.Fullname, &dat.Nickname, &dat.About)
		if err != nil {
			return nil, err
		}
		users = append(users, &dat)
	}
	return users, nil
}

func (s *Store) GetProfile(nickname string) (*model.User, error) {
	rows, err := s.db.Query(context.Background(), `SELECT email, fullname, nickname, about FROM users WHERE nickname = $1;`, nickname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		dat := model.User{}
		err := rows.Scan(&dat.Email, &dat.Fullname, &dat.Nickname, &dat.About)
		if err != nil {
			return nil, err
		}
		return &dat, nil
	}
	return nil, model.ErrNotFound404
}

func (s *Store) ChangeProfile(in *model.User) error {
	_, err := s.db.Exec(context.Background(), `UPDATE users SET email = $1, fullname = $2, about = $3 WHERE nickname = $4;`, in.Email, in.Fullname, in.About, in.Nickname)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateForum(in *model.Forum) error {
	_, err := s.db.Exec(context.Background(), `INSERT INTO forums (title, user1, slug) VALUES ($1, $2, $3);`, in.Title, in.User, in.Slug)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) GetForumByUsername(nickname string) (*model.Forum, error) {
	rows, err := s.db.Query(context.Background(), `SELECT title, user1, slug, posts, threads FROM forums WHERE user1 = $1;`, nickname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		dat := model.Forum{}
		err := rows.Scan(&dat.Title, &dat.User, &dat.Slug, &dat.Posts, &dat.Threads)
		if err != nil {
			return nil, err
		}
		return &dat, nil
	}
	return nil, model.ErrNotFound404
}

func (s *Store) GetForumBySlug(slug string) (*model.Forum, error) {
	rows, err := s.db.Query(context.Background(), `SELECT title, user1, slug, posts, threads FROM forums WHERE slug = $1;`, slug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		dat := model.Forum{}
		err := rows.Scan(&dat.Title, &dat.User, &dat.Slug, &dat.Posts, &dat.Threads)
		if err != nil {
			return nil, err
		}
		return &dat, nil
	}
	return nil, model.ErrNotFound404
}

func (s *Store) GetThreadByModel(in *model.Thread) (*model.Thread, error) {
	rows, err := s.db.Query(context.Background(), `SELECT title, author, forum, message, votes, slug, created FROM threads WHERE title = $1 AND author = $2 AND message = $3;`, in.Title, in.Author, in.Message)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		dat := model.Thread{}
		err := rows.Scan(&dat.Title, &dat.Author, &dat.Forum, &dat.Message, &dat.Votes, &dat.Slug, &dat.Created)
		if err != nil {
			return nil, err
		}
		return &dat, nil
	}
	return nil, model.ErrNotFound404
}

func (s *Store) CreateThreadByModel(in *model.Thread) (*model.Thread, error) {
	createTime := time.Now()
	_, err := s.db.Exec(context.Background(), `INSERT INTO threads (title, author, forum, message, votes, slug, created) VALUES ($1, $2, $3, $4, $5, $6, $7);`, in.Title, in.Author, in.Forum, in.Message, 0, in.Slug, createTime.Format("2006.01.02 15:04:05"))
	if err != nil {
		return nil, err
	}
	in.Created = createTime
	return in, nil
}

func (s *Store) GetForumUsers(slug string, limit int, since string, desc bool) ([]*model.User, error) {
	users := []*model.User{}
	rows, err := s.db.Query(context.Background(), `SELECT * FROM (SELECT email, fullname, nickname, about FROM users JOIN posts ON users.nickname=posts.author WHERE posts.forum = $1
		UNION SELECT email, fullname, nickname, about FROM users JOIN threads ON users.nickname=threads.author WHERE threads.forum = $1) AS U WHERE U.nickname > $2 ORDER BY U.nickname LIMIT $3;`, slug, since, limit)
	if err != nil {
		return nil, err
	}
	if desc {
		if since == "" {
			since = "яяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяя"
		}
		rows, err = s.db.Query(context.Background(), `SELECT * FROM (SELECT email, fullname, nickname, about FROM users JOIN posts ON users.nickname=posts.author WHERE posts.forum = $1
			UNION SELECT email, fullname, nickname, about FROM users JOIN threads ON users.nickname=threads.author WHERE threads.forum = $1) AS U WHERE U.nickname < $2 ORDER BY U.nickname DESC LIMIT $3;`, slug, since, limit)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		dat := model.User{}
		err := rows.Scan(&dat.Email, &dat.Fullname, &dat.Nickname, &dat.About)
		if err != nil {
			return nil, err
		}
		users = append(users, &dat)
	}
	return users, nil
}
