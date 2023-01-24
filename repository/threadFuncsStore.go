package repository

import (
	"dbproject/model"
	"time"
)

func (s *Store) CheckAllPostParentIds(in []int) error {
	dbcount := 0
	err := s.db.QueryRow(`SELECT count(*) FROM (SELECT parent FROM posts JOIN threads ON posts.thread = threads.id WHERE threads.id = 1 GROUP BY parent HAVING $1 @> array_agg(parent)) AS S;`, in).Scan(&dbcount)
	if err != nil {
		return err
	}
	if dbcount < len(in) {
		return model.ErrConflict409
	}
	return nil
}

func (s *Store) GetThreadById(id int) (*model.Thread, error) {
	rows, err := s.db.Query(`SELECT id, title, author, forum, message, votes, slug, created FROM threads WHERE id = $1;`, id)
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
	rows, err := s.db.Query(`SELECT id, title, author, forum, message, votes, slug, created FROM threads WHERE LOWER(slug) = LOWER($1);`, slug)
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
	createdFormatted := createTime.Format(time.RFC3339)
	for _, post := range *in {
		id := 0
		insertModel := model.Post{Parent: post.Parent, Author: post.Author, Message: post.Message, IsEdited: false, Thread: threadId, Forum: forumSlug, Created: createTime}
		err := s.db.QueryRow(`INSERT INTO posts (parent, author, message, forum, thread, isedited, created) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`, insertModel.Parent, insertModel.Author, insertModel.Message, insertModel.Forum, insertModel.Thread, insertModel.IsEdited, createdFormatted).Scan(&id)
		if err != nil {
			return nil, err
		}
		insertModel.Id = id
		posts = append(posts, &insertModel)
	}
	_, err := s.db.Exec(`UPDATE forums SET posts = posts + $1 WHERE slug = $2;`, len(*in), forumSlug)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *Store) UpdateThreadInfo(in *model.ThreadUpdate, id int) error {
	_, err := s.db.Exec(`UPDATE threads SET message = $1, title = $2 WHERE id = $3;`, in.Message, in.Title, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) VoteForThread(in *model.Vote, threadID int) (int, error) {
	count := 0
	err := s.db.QueryRow(`SELECT count(*) FROM (SELECT users.nickname FROM users JOIN votes on users.nickname=votes.nickname JOIN threads ON votes.thread=threads.id WHERE threads.id = $1) AS S WHERE S.nickname IN ($2);`, threadID, in.Nickname).Scan(&count)
	if err != nil {
		return 0, err
	}
	if count == 0 {
		_, err := s.db.Exec(`INSERT INTO votes (nickname, thread, voice) VALUES ($1, $2, $3);`, in.Nickname, threadID, in.Voice)
		if err != nil {
			return 0, err
		}
	} else {
		_, err := s.db.Exec(`UPDATE votes SET voice = $1 WHERE nickname = $2 AND thread = $3;`, in.Voice, in.Nickname, threadID)
		if err != nil {
			return 0, err
		}
	}
	newRate := 0
	err = s.db.QueryRow(`SELECT sum(voice) FROM votes WHERE thread = $1;`, threadID).Scan(&newRate)
	if err != nil {
		return 0, err
	}
	_, err = s.db.Exec(`UPDATE threads SET votes = $1 WHERE id = $2;`, newRate, threadID)
	if err != nil {
		return 0, err
	}
	return newRate, nil
}

func (s *Store) GetThreadPosts(threadId int, limit int, since int, sort string, desc bool) ([]*model.Post, error) {
	posts := []*model.Post{}

	rows, err := s.db.Query(`SELECT posts.id, parent, posts.author, posts.message, posts.forum, posts.thread, isedited, posts.created FROM posts JOIN threads ON posts.thread = threads.id WHERE threads.id = $1 AND posts.id > $2 ORDER BY posts.created LIMIT $3;`, threadId, since, limit)
	if err != nil {
		return nil, err
	}
	if desc {
		if since == 0 {
			since = 1e9
		}
		rows, err = s.db.Query(`SELECT posts.id, parent, posts.author, posts.message, posts.forum, posts.thread, isedited, posts.created FROM posts JOIN threads ON posts.thread = threads.id WHERE threads.id = $1 AND posts.id < $2 ORDER BY posts.created DESC LIMIT $3;`, threadId, since, limit)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		dat := model.Post{}
		err := rows.Scan(&dat.Id, &dat.Parent, &dat.Author, &dat.Message, &dat.Forum, &dat.Thread, &dat.IsEdited, &dat.Created)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &dat)
	}
	return posts, nil
}
