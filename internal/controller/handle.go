package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"user-service/internal/core/entity/Error_Code"
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
		if r.Method == "GET" {
			http.ServeFile(w, r, "assets/student/views/SignUp.html")
		} else if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(reqBody))
			var user request.SignUpRequest
			json.Unmarshal(reqBody, &user)

			var rep *response.Response = u.userService.SignUp(user)

			if rep.Error_code == Error_Code.InvalidRequest {
				// Nguoi dung chua nhap username hoac mat khau
				w.WriteHeader(http.StatusUnauthorized)
			} else if rep.Error_code == Error_Code.DuplicateUser {
				// Username nguoi dung nhap bi da co roi
				w.WriteHeader(http.StatusConflict)
			}

			repJson, _ := json.Marshal(rep)
			w.Header().Add("Content-Type", "application/json")
			w.Write(repJson)
		}
	}
}

func ReturnLoginStudentPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/student/views/Login.html")
}
