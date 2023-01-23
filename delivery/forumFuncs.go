package delivery

import (
	"dbproject/model"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// CreateForum godoc
// @Summary Creates Forum
// @Description Creates Forum
// @ID CreateForum
// @Accept  json
// @Produce  json
// @Tags Forum
// @Param forum body model.ForumCreateModel true "Forum params"
// @Success 201 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 404 {object} model.Error "Not found - Requested entity is not found in database"
// @Failure 409 {object} model.Error "Conflict - User already exists"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /forum/create [post]
func (api *Handler) CreateForum(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req model.Forum
	err := decoder.Decode(&req)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, model.ErrBadRequest400, 400)
		return
	}
	_, err = api.usecase.GetProfile(req.User)
	if err == model.ErrNotFound404 {
		log.Println(err)
		ReturnErrorJSON(w, model.ErrNotFound404, 404)
		return
	}
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, model.ErrServerError500, 500)
		return
	}

	forum, err := api.usecase.GetForumByUsername(req.User)
	if err != nil && err != model.ErrNotFound404 {
		log.Println(err)
		ReturnErrorJSON(w, model.ErrServerError500, 500)
		return
	}
	if forum != nil {
		w.WriteHeader(409)
		json.NewEncoder(w).Encode(&forum)
		return
	}
	err = api.usecase.CreateForum(&req)
	if err != nil {
		log.Println("db err: ", err)
		ReturnErrorJSON(w, model.ErrServerError500, 500)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&req)
}

// GetForumInfo godoc
// @Summary Gets forum info
// @Description Gets forum info
// @ID GetForumInfo
// @Accept  json
// @Produce  json
// @Tags Forum
// @Param slug path string true "slug of user"
// @Success 200 {object} model.Forum
// @Failure 404 {object} model.Error "Not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /forum/{slug}/details [get]
func (api *Handler) GetForumInfo(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Path, "/")
	slug := s[len(s)-2]

	forum, err := api.usecase.GetForumBySlug(slug)
	if err == model.ErrNotFound404 {
		log.Println(err)
		ReturnErrorJSON(w, model.ErrNotFound404, 404)
		return
	}
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, model.ErrServerError500, 500)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&forum)
}

// CreateThread godoc
// @Summary creates thread
// @Description creates thread
// @ID CreateThread
// @Accept  json
// @Produce  json
// @Tags Forum
// @Param slug path string true "slug of thread"
// @Param thread body model.ThreadCreateModel true "Thread params"
// @Success 201 {object} model.Thread
// @Failure 404 {object} model.Error "Not found - Requested entity is not found in database"
// @Failure 409 {object} model.Thread
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /forum/{slug}/create [post]
func (api *Handler) CreateThread(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Path, "/")
	slug := s[len(s)-2]
	var req model.Thread
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, model.ErrBadRequest400, 400)
		return
	}
	req.Slug = slug
	req.Votes = 0

	forum, err := api.usecase.GetForumBySlug(slug)
	if err == model.ErrNotFound404 {
		log.Println(err)
		ReturnErrorJSON(w, model.ErrNotFound404, 404)
		return
	}
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, model.ErrServerError500, 500)
		return
	}

	req.Forum = forum.Title

	_, err = api.usecase.GetProfile(req.Author)
	if err == model.ErrNotFound404 {
		log.Println(err)
		ReturnErrorJSON(w, model.ErrNotFound404, 404)
		return
	}
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, model.ErrServerError500, 500)
		return
	}

	thread, err := api.usecase.GetThreadByModel(&req)
	if err != nil && err != model.ErrNotFound404 {
		log.Println(err)
		ReturnErrorJSON(w, model.ErrServerError500, 500)
		return
	}
	if thread != nil {
		w.WriteHeader(409)
		json.NewEncoder(w).Encode(&thread)
		return
	}
	thread, err = api.usecase.CreateThreadByModel(&req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, model.ErrServerError500, 500)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&thread)
}
