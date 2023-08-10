package controller

import (
	"net/http"
)

type UserController struct {
	Mux *http.ServeMux
}

func NewUserController() *UserController {
	return &UserController{
		Mux: http.NewServeMux(),
	}
}

func (u UserController) Router() {
	u.Mux.HandleFunc("/", HandleMain)
}
