package app

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFindNoUser(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/user/99999")
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("http.Get error %v", err)
	}

	badRequest := 400
	if resp.StatusCode != badRequest {
		t.Errorf("[StatusCode] got: %v, want: %v", resp.StatusCode, badRequest)
	}

	json := "{\"error\":\"Bad Request\"}"
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll error %v", err)
	}
	strBody := string(body)
	if strBody != json {
		t.Errorf("[Body] got: %v, want: %v", strBody, json)
	}
}
