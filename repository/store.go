package repository

import (
	"dbproject/model"
	"time"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

type StoreInterface interface {
	CreateUser(params *model.User) error
	GetUsersByUsermodel(in *model.User) ([]*model.User, error)
	GetProfile(nickname string) (*model.User, error)
	ChangeProfile(in *model.User) error
	CreateForum(in *model.Forum) error
	GetForumByUsername(nickname string) (*model.Forum, error)
	GetForumBySlug(slug string) (*model.Forum, error)
	CreateThreadByModel(in *model.Thread) (*model.Thread, error)
	GetForumUsers(slug string, limit int, since string, desc bool) ([]*model.User, error)
	GetForumThreads(slug string, limit int, since time.Time, desc bool) ([]*model.Thread, error)
	GetPostById(id int) (*model.Post, error)
	UpdatePost(id int, mes string) (*model.Post, error)
	GetServiceStatus() (*model.Status, error)
	ServiceClear() error
	CheckAllPostParentIds(threadId int, in []int) error
	GetThreadById(id int) (*model.Thread, error)
	GetThreadBySlug(slug string) (*model.Thread, error)
	CreatePosts(in *model.Posts, threadId int, forumSlug string) ([]*model.Post, error)
	UpdateThreadInfo(in *model.ThreadUpdate, id int) error
	VoteForThread(in *model.Vote, threadID int) (int, error)
	GetThreadPostsFlatSort(threadId int, limit int, since int, desc bool) ([]*model.Post, error)
	GetThreadPostsTreeSort(threadId int, limit int, since int, desc bool) ([]*model.Post, error)
}

type Store struct {
	db *pgx.ConnPool
}

func NewStore(db *pgx.ConnPool) StoreInterface {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateUser(in *model.User) error {
	_, err := s.db.Exec(`INSERT INTO users (email, fullname, nickname, about) VALUES ($1, $2, $3, $4);`, in.Email, in.Fullname, in.Nickname, in.About)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) GetUsersByUsermodel(in *model.User) ([]*model.User, error) {
	users := []*model.User{}
	rows, err := s.db.Query(`SELECT email, fullname, nickname, about FROM users WHERE LOWER(email) = LOWER($1) OR LOWER(nickname) = LOWER($2);`, in.Email, in.Nickname)
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
	rows, err := s.db.Query(`SELECT email, fullname, nickname, about FROM users WHERE LOWER(nickname) = LOWER($1);`, nickname)
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
	_, err := s.db.Exec(`UPDATE users SET email = $1, fullname = $2, about = $3 WHERE LOWER(nickname) = LOWER($4);`, in.Email, in.Fullname, in.About, in.Nickname)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateForum(in *model.Forum) error {
	_, err := s.db.Exec(`INSERT INTO forums (title, user1, slug) VALUES ($1, $2, $3);`, in.Title, in.User, in.Slug)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) GetForumByUsername(nickname string) (*model.Forum, error) {
	rows, err := s.db.Query(`SELECT title, user1, slug, posts, threads FROM forums WHERE LOWER(user1) = LOWER($1);`, nickname)
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
	rows, err := s.db.Query(`SELECT title, user1, slug, posts, threads FROM forums WHERE LOWER(slug) = LOWER($1);`, slug)
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

func (s *Store) CreateThreadByModel(in *model.Thread) (*model.Thread, error) {
	id := 0
	err := s.db.QueryRow(`INSERT INTO threads (title, author, forum, message, votes, slug, created) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`, in.Title, in.Author, in.Forum, in.Message, 0, in.Slug, in.Created).Scan(&id)
	if err != nil {
		return nil, err
	}
	_, err = s.db.Exec(`UPDATE forums SET threads = threads + 1 WHERE LOWER(slug) = LOWER($1);`, in.Forum)
	if err != nil {
		return nil, err
	}
	in.Id = id
	return in, nil
}

func (s *Store) GetForumUsers(slug string, limit int, since string, desc bool) ([]*model.User, error) {
	users := []*model.User{}
	rows, err := s.db.Query(`SELECT * FROM (SELECT email, fullname, nickname, about FROM users JOIN posts ON LOWER(users.nickname)=LOWER(posts.author) WHERE LOWER(posts.forum) = LOWER($1)
		UNION SELECT email, fullname, nickname, about FROM users JOIN threads ON LOWER(users.nickname)=LOWER(threads.author) WHERE LOWER(threads.forum) = LOWER($1)) AS U WHERE LOWER(U.nickname) > LOWER($2) ORDER BY LOWER(U.nickname) ASC LIMIT $3;`, slug, since, limit)
	if err != nil {
		return nil, err
	}
	if desc {
		if since == "" {
			since = "яяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяя"
		}
		rows, err = s.db.Query(`SELECT * FROM (SELECT email, fullname, nickname, about FROM users JOIN posts ON LOWER(users.nickname)=LOWER(posts.author) WHERE LOWER(posts.forum) = LOWER($1)
		UNION SELECT email, fullname, nickname, about FROM users JOIN threads ON LOWER(users.nickname)=LOWER(threads.author) WHERE LOWER(threads.forum) = LOWER($1)) AS U WHERE LOWER(U.nickname) < LOWER($2) ORDER BY LOWER(U.nickname) DESC LIMIT $3;`, slug, since, limit)
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

func (s *Store) GetForumThreads(slug string, limit int, since time.Time, desc bool) ([]*model.Thread, error) {
	threads := []*model.Thread{}

	rows, err := s.db.Query(`SELECT id, title, author, forum, message, votes, slug, created FROM threads WHERE LOWER(forum) = LOWER($1) AND created >= $2 ORDER BY created ASC LIMIT $3;`, slug, since, limit)
	if err != nil {
		return nil, err
	}
	if desc {
		if since.Format("0000-01-01T00:00:00.000Z") == "0000-01-01T00:00:00.000Z" {
			since, err = time.Parse(time.RFC3339, "5000-01-01T00:00:00.000Z")
			if err != nil {
				return nil, err
			}
		}
		rows, err = s.db.Query(`SELECT id, title, author, forum, message, votes, slug, created FROM threads WHERE LOWER(forum) = LOWER($1) AND created <= $2 ORDER BY created DESC LIMIT $3;`, slug, since, limit)
	}
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
		threads = append(threads, &dat)
	}
	return threads, nil
}

func (s *Store) GetPostById(id int) (*model.Post, error) {
	rows, err := s.db.Query(`SELECT id, parent, author, message, forum, isedited, thread, created FROM posts WHERE id = $1;`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		dat := model.Post{}
		err := rows.Scan(&dat.Id, &dat.Parent, &dat.Author, &dat.Message, &dat.Forum, &dat.IsEdited, &dat.Thread, &dat.Created)
		if err != nil {
			return nil, err
		}
		return &dat, nil
	}
	return nil, model.ErrNotFound404
}

func (s *Store) UpdatePost(id int, mes string) (*model.Post, error) {
	post, err := s.GetPostById(id)
	if err != nil {
		return nil, err
	}
	if post.Message == mes || mes == "" {
		return post, nil
	}
	_, err = s.db.Exec(`UPDATE posts SET message = $1, isedited = $2 WHERE id = $3;`, mes, true, id)
	if err != nil {
		return nil, err
	}
	post.Message = mes
	post.IsEdited = true
	return post, nil
}

func (s *Store) GetServiceStatus() (*model.Status, error) {
	status := &model.Status{}
	rows, err := s.db.Query(`SELECT (SELECT count(*) FROM users) AS u, (SELECT count(*) FROM forums) AS f, (SELECT count(*) FROM threads) AS t, (SELECT count(*) FROM posts) AS p;`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&status.User, &status.Forum, &status.Thread, &status.Post)
		if err != nil {
			return nil, err
		}
	}
	return status, nil
}

func (s *Store) ServiceClear() error {
	_, err := s.db.Exec("TRUNCATE TABLE users, forums, threads, posts, votes CASCADE;")
	if err != nil {
		return err
	}
	return nil
}
