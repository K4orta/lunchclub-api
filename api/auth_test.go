package api

import "testing"

func TestFBLoginURL(t *testing.T) {
	fbClientID = "123"
	stubRedirect := "http://testurl.com/auth"
	url := FBLoginURL(stubRedirect)
	testURL := `https://www.facebook.com/dialog/oauth?access_type=online&client_id=123&redirect_uri=http%3A%2F%2Ftesturl.com%2Fauth&response_type=code&scope=public_profile+email&state=state`
	if url != testURL {
		t.Errorf("Did format auth URL correctly, expected: %v got: %v", testURL, url)
	}
}
