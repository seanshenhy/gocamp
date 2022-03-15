package main

import (
	"gocamp/ch1/ctrl"
	"gocamp/ch1/model"
	"net/http"
)

func main() {
	model.InitDb()
	http.HandleFunc("/userinfo", ctrl.UserList)
	http.ListenAndServe(":8080", nil)

}
