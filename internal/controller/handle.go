package controller

import (
	"net/http"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to My Website </h1>"))
}
