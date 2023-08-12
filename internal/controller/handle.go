package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/index.html")
}

func ReturnSignUpStudentPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/student/views/SignUp.html")
}

func SignUp(u UserController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		var user request.SignUpRequest
		json.Unmarshal(reqBody, &user)

		var rep *response.Response = u.userService.SignUp(user)

		repJson, _ := json.Marshal(rep)
		fmt.Println(string(repJson))
		w.Write(repJson)
	}
}

func ReturnLoginStudentPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ok")
	http.ServeFile(w, r, "assets/student/views/Login.html")
}
