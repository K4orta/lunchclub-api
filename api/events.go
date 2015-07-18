package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type eventForm struct {
	YelpURL   string
	StartTime string
}

// CreateEvent handles a POST request to create a new Event and save it in the DB
func CreateEvent(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	var ef eventForm
	err = json.Unmarshal(b, &ef)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	// db, _ := storage.CreateConnection()
	// defer db.Close()
	fmt.Fprint(w, string(b))
	// loc, err := storage.GetLocationBySlug(db, "")
}

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
