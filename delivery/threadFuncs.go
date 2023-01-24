package delivery

import (
	"dbproject/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// CreatePosts godoc
// @Summary Creates Posts
// @Description Creates Posts
// @ID CreatePosts
// @Accept  json
// @Produce  json
// @Tags Thread
// @Param posts body model.Posts true "Posts params"
// @Param slug_or_id path string true "slug or id"
// @Success 201 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 404 {object} model.Error "Not found - Requested entity is not found in database"
// @Failure 409 {object} model.Error "Conflict - User already exists"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /thread/{slug_or_id}/create [post]
func (api *Handler) CreatePosts(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Path, "/")
	slug_or_id := s[len(s)-2]
	id := 0
	slug := slug_or_id
	id, err := strconv.Atoi(slug_or_id)
	if err != nil {
		log.Println("error: ", err)

	}

	decoder := json.NewDecoder(r.Body)
	var req model.Posts
	err = decoder.Decode(&req)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, model.ErrBadRequest400, 400)
		return
	}
	posts, err := api.usecase.CreatePosts(&req, id, slug)
	if err == model.ErrNotFound404 {
		log.Println(err)
		ReturnErrorJSON(w, model.ErrNotFound404, 404)
		return
	}
	if err == model.ErrConflict409 {
		log.Println(err)
		ReturnErrorJSON(w, model.ErrConflict409, 409)
		return
	}
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, model.ErrServerError500, 500)
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&posts)
}
