package ctrl

import (
	"encoding/json"
	"fmt"
	"gocamp/ch1/service"
	"net/http"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	userInfo, err := service.GetUserInfo(1)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	data, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	fmt.Fprintf(w, string(data))
}
