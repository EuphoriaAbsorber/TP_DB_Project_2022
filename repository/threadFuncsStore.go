package repository

import (
	"dbproject/model"
	"time"

	"github.com/jackc/pgx"
)

func (s *Store) CheckAllPostParentIds(threadId int, in []int) error {
	dbcount := 0
	err := s.db.QueryRow(`SELECT count(*) FROM (SELECT id FROM posts WHERE thread = $1 GROUP BY id HAVING $2 @> array_agg(id)) AS S;`, threadId, in).Scan(&dbcount)
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
	if slug == "" {
		return nil, nil
	}
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
	dbCreatedTime := time.Now()
	for _, post := range *in {
		id := 0
		insertModel := model.Post{Parent: post.Parent, Author: post.Author, Message: post.Message, IsEdited: false, Thread: threadId, Forum: forumSlug, Created: createTime}
		err := s.db.QueryRow(`INSERT INTO posts (parent, author, message, forum, thread, isedited, created) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created;`, insertModel.Parent, insertModel.Author, insertModel.Message, insertModel.Forum, insertModel.Thread, insertModel.IsEdited, createdFormatted).Scan(&id, &dbCreatedTime)
		if err != nil {
			return nil, err
		}
		insertModel.Id = id
		insertModel.Created = dbCreatedTime
		posts = append(posts, &insertModel)
	}
	_, err := s.db.Exec(`UPDATE forums SET posts = posts + $1 WHERE LOWER(slug) = LOWER($2);`, len(*in), forumSlug)
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
	err := s.db.QueryRow(`SELECT count(*) FROM votes WHERE thread = $1 AND nickname = $2;`, threadID, in.Nickname).Scan(&count)
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

func (s *Store) GetThreadPostsFlatSort(threadId int, limit int, since int, desc bool) ([]*model.Post, error) {
	posts := []*model.Post{}

	rows, err := s.db.Query(`SELECT id, parent, author, message, forum, thread, isedited, created FROM posts WHERE thread = $1 AND id > $2 ORDER BY (created, id) LIMIT $3;`, threadId, since, limit)
	if err != nil {
		return nil, err
	}
	if desc {
		if since == 0 {
			since = 1e9
		}
		rows, err = s.db.Query(`SELECT id, parent, author, message, forum, thread, isedited, created FROM posts WHERE thread = $1 AND id < $2 ORDER BY (created, id) DESC LIMIT $3;`, threadId, since, limit)
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

func (s *Store) GetThreadPostsTreeSort(threadId int, limit int, since int, desc bool) ([]*model.Post, error) {
	posts := []*model.Post{}
	var rows *pgx.Rows
	var err error

	if since == 0 {
		rows, err = s.db.Query(`SELECT id, COALESCE(parent, 0), author, message, forum, thread, isedited, created FROM posts WHERE thread = $1 ORDER BY path LIMIT $2;`, threadId, limit)
	} else {
		rows, err = s.db.Query(`SELECT id, COALESCE(parent, 0), author, message, forum, thread, isedited, created FROM posts WHERE thread = $1 AND path > (SELECT path FROM posts WHERE id = $2) ORDER BY path LIMIT $3;`, threadId, since, limit)
	}
	if desc {
		if since == 0 {
			rows, err = s.db.Query(`SELECT id, COALESCE(parent, 0), author, message, forum, thread, isedited, created FROM posts WHERE thread = $1 ORDER BY path DESC LIMIT $2;`, threadId, limit)
		} else {
			rows, err = s.db.Query(`SELECT id, COALESCE(parent, 0), author, message, forum, thread, isedited, created FROM posts WHERE thread = $1 AND path < (SELECT path FROM posts WHERE id = $2) ORDER BY path DESC LIMIT $3;`, threadId, since, limit)
		}
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
