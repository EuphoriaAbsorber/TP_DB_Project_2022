package delivery

import (
	"dbproject/model"
	usecase "dbproject/usecase"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// @title DB project API
// @version 1.0
// @description DB project server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath  /api

type Handler struct {
	usecase usecase.UsecaseInterface
}

func NewHandler(uc usecase.UsecaseInterface) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func ReturnErrorJSON(w http.ResponseWriter, err error, errCode int) {
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&model.Error{Error: err.Error()})
	return
}

// CreateUser godoc
// @Summary Creates User
// @Description Creates User
// @ID CreateUser
// @Accept  json
// @Produce  json
// @Tags User
// @Param nickname path string true "nickname of user"
// @Param user body model.User true "User params"
// @Success 201 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 409 {object} model.Error "Conflict - User already exists"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /user/{nickname}/create [post]
func (api *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Path, "/")
	nickname := s[len(s)-2]
	decoder := json.NewDecoder(r.Body)
	var req model.User
	err := decoder.Decode(&req)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, model.ErrBadRequest400, 400)
		return
	}
	req.Nickname = nickname

	users, err := api.usecase.GetUsersByUsermodel(&req)
	if err != nil {
		log.Println("get GetUserByUsermodel ", err)
		ReturnErrorJSON(w, model.ErrServerError500, 500)
		return
	}

	if len(users) > 0 {
		w.WriteHeader(409)
		json.NewEncoder(w).Encode(&users)
		return
	}

	err = api.usecase.CreateUser(&req)
	if err != nil {
		log.Println("db err: ", err)
		ReturnErrorJSON(w, model.ErrServerError500, 500)
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&req)
}
