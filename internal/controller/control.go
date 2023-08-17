package controller

import (
	"net/http"
	"user-service/internal/core/port/service"
)

type UserController struct {
	Mux         *http.ServeMux
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		Mux:         http.NewServeMux(),
		userService: userService,
	}
}

func (u UserController) Router() {
	u.Mux.HandleFunc("/", HandleMain)

	fs := http.FileServer(http.Dir("assets"))
	u.Mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//u.Mux.HandleFunc("/student/signup", ReturnSignUpStudentPage)

	u.Mux.HandleFunc("/student/signup", SignUp(u))
	u.Mux.HandleFunc("/student/login", ReturnLoginStudentPage)
}
