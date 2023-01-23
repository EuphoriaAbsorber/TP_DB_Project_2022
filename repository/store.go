package repository

import (
	"context"
	"dbproject/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StoreInterface interface {
	CreateUser(params *model.User) error
	GetUsersByUsermodel(in *model.User) ([]*model.User, error)
	GetProfile(nickname string) (*model.User, error)
	ChangeProfile(in *model.User) error
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
