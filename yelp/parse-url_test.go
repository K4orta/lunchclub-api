package yelp

import "testing"

// TestParseURL tests the expected case for transforming a URL to Yelp ID
func TestParseURL(t *testing.T) {
	example := "http://www.yelp.com/biz/kona-club-oakland"
	if ParseURL(example) != "kona-club-oakland" {
		t.Errorf("Error parsing URL. Expected: %v got: %v", "kona-club-oakland", ParseURL(example))
	}
}
