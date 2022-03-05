package main

import (
	"gocamp/ch1/ctrl"
	"net/http"

	"gocamp/ch1/model"
)

func main() {
	model.InitDb()
	http.HandleFunc("/userinfo", ctrl.UserList)
	http.ListenAndServe(":8080", nil)
}
