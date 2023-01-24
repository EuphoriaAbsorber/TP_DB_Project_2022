package repository

import (
	"context"
	"dbproject/model"
	"time"
)

func (s *Store) CheckAllPostParentIds(in []int) error {
	dbcount := 0
	err := s.db.QueryRow(context.Background(), `SELECT count(*) FROM (SELECT parent FROM posts JOIN threads ON posts.thread = threads.id WHERE threads.id = 1 GROUP BY parent HAVING $1 @> array_agg(parent)) AS S;`, in).Scan(&dbcount)
	if err != nil {
		return err
	}
	if dbcount < len(in) {
		return model.ErrConflict409
	}
	return nil
}

func (s *Store) GetThreadById(id int) (*model.Thread, error) {
	rows, err := s.db.Query(context.Background(), `SELECT id, title, author, forum, message, votes, slug, created FROM threads WHERE id = $1;`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		dat := model.Thread{}
		err := rows.Scan(&dat.Id, &dat.Title, &dat.Author, &dat.Forum, &dat.Message, &dat.Votes, &dat.Slug, &dat.Created)
		if err != nil {
			return nil, err
		}
		return &dat, nil
	}
	return nil, model.ErrNotFound404
}

func (s *Store) GetThreadBySlug(slug string) (*model.Thread, error) {
	rows, err := s.db.Query(context.Background(), `SELECT id, title, author, forum, message, votes, slug, created FROM threads WHERE slug = $1;`, slug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		dat := model.Thread{}
		err := rows.Scan(&dat.Id, &dat.Title, &dat.Author, &dat.Forum, &dat.Message, &dat.Votes, &dat.Slug, &dat.Created)
		if err != nil {
			return nil, err
		}
		return &dat, nil
	}
	return nil, model.ErrNotFound404
}

func (s *Store) CreatePosts(in *model.Posts, threadId int, forumSlug string) ([]*model.Post, error) {

	posts := []*model.Post{}
	createTime := time.Now()
	for _, post := range in.Posts {
		id := 0
		insertModel := model.Post{Parent: post.Parent, Author: post.Author, Message: post.Message, IsEdited: false, Thread: threadId, Forum: forumSlug, Created: createTime}
		err := s.db.QueryRow(context.Background(), `INSERT INTO posts (parent, author, message, forum, thread, isedited, created) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`, insertModel.Parent, insertModel.Author, insertModel.Message, insertModel.Forum, insertModel.Thread, insertModel.IsEdited, createTime.Format("2006.01.02 15:04:05")).Scan(&id)
		if err != nil {
			return nil, err
		}
		insertModel.Id = id
		posts = append(posts, &insertModel)
	}
	_, err := s.db.Exec(context.Background(), `UPDATE forums SET posts = posts + $1 WHERE slug = $2;`, len(in.Posts), forumSlug)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *Store) UpdateThreadInfo(in *model.ThreadUpdate) error {
	_, err := s.db.Exec(context.Background(), `UPDATE threads SET message = $1, title = $2;`, in.Message, in.Title)
	if err != nil {
		return err
	}
	return nil
}