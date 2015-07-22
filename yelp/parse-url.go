package yelp

import "strings"

// ParseURL takes a yelp url and return the business ID
func ParseURL(url string) string {
	parts := strings.Split(url, "/biz/")
	return parts[1]
}
