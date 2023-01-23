package delivery

import (
	"dbproject/model"
	usecase "dbproject/usecase"
	"encoding/json"
	"net/http"
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
