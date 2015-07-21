package api

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCreateEvent(t *testing.T) {
	fs := mockAPI(HelloWorld)
	resp, _ := http.Get(fs.URL)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	if string(b) != "Hello World" {
		t.Errorf("Expected 'Hello World' got %v", string(b))
	}
}
