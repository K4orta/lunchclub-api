package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, req *http.Request) {

}

func ReadUsers(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// users, err := storage.FetchAllUsers()
	// if err != nil {
	// 	sendError(w, "Error")
	// 	return
	// }

	users := []string{"abc", "dfg"}

	out, err := json.Marshal(users)
	if err == nil {
		fmt.Fprint(w, string(out))
	} else {
		sendError(w, "Error")
	}
}
