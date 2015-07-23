package yelp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchBusiness(t *testing.T) {
	fs := makeFakeServer(stubBusinessResponse())
	businessEndpoint = fs.URL + "/"

	b, err := FetchBusiness("kona-club-oakland")
	if err != nil {
		t.Error(err)
	}

	if b.Slug != "kona-club-oakland" {
		t.Error("expected business slug to equal 'kona-club-oakland', got %v", b.Slug)
	}
}

func stubBusinessResponse() string {
	return `{
    "is_claimed": true,
    "rating": 4.0,
    "mobile_url": "http://m.yelp.com/biz/kona-club-oakland",
    "rating_img_url": "http://s3-media4.fl.yelpcdn.com/assets/2/www/img/c2f3dd9799a5/ico/stars/v1/stars_4.png",
    "review_count": 430,
    "name": "Kona Club",
    "rating_img_url_small": "http://s3-media4.fl.yelpcdn.com/assets/2/www/img/f62a5be2f902/ico/stars/v1/stars_small_4.png",
    "url": "http://www.yelp.com/biz/kona-club-oakland",
    "is_closed": false,
    "reviews": [
        {
            "rating": 5,
            "excerpt": "Kona Club, sister bar to the Mallard, is a perfect guilty pleasure. It's comfortable, the bartenders are nice, and the drinks are tasty. Scorpion bowls are...",
            "time_created": 1435031069,
            "rating_image_url": "http://s3-media1.fl.yelpcdn.com/assets/2/www/img/f1def11e4e79/ico/stars/v1/stars_5.png",
            "rating_image_small_url": "http://s3-media1.fl.yelpcdn.com/assets/2/www/img/c7623205d5cd/ico/stars/v1/stars_small_5.png",
            "user": {
                "image_url": "http://s3-media3.fl.yelpcdn.com/photo/OC00DOtdGvZVu-i6jQGMQw/ms.jpg",
                "id": "6ZEtkXaGj10rQDgOtfZFsw",
                "name": "Kellen F."
            },
            "rating_image_large_url": "http://s3-media3.fl.yelpcdn.com/assets/2/www/img/22affc4e6c38/ico/stars/v1/stars_large_5.png",
            "id": "eUxXCRHCprsNdx6Zub6c5Q"
        }
    ],
    "phone": "5106547100",
    "snippet_text": "Kona Club, sister bar to the Mallard, is a perfect guilty pleasure. It's comfortable, the bartenders are nice, and the drinks are tasty. Scorpion bowls are...",
    "image_url": "http://s3-media3.fl.yelpcdn.com/bphoto/f_hzJrqUW86uBgzPp02tgg/ms.jpg",
    "categories": [
        [
            "Lounges",
            "lounges"
        ]
    ],
    "display_phone": "+1-510-654-7100",
    "rating_img_url_large": "http://s3-media2.fl.yelpcdn.com/assets/2/www/img/ccf2b76faa2c/ico/stars/v1/stars_large_4.png",
    "id": "kona-club-oakland",
    "snippet_image_url": "http://s3-media3.fl.yelpcdn.com/photo/OC00DOtdGvZVu-i6jQGMQw/ms.jpg",
    "location": {
        "cross_streets": "Pleasant Valley Ave & Ramona Ave",
        "city": "Oakland",
        "display_address": [
            "4401 Piedmont Ave",
            "North Oakland",
            "Oakland, CA 94611"
        ],
        "geo_accuracy": 8.0,
        "neighborhoods": [
            "North Oakland",
            "Piedmont Ave"
        ],
        "postal_code": "94611",
        "country_code": "US",
        "address": [
            "4401 Piedmont Ave"
        ],
        "coordinate": {
            "latitude": 37.830536,
            "longitude": -122.247383
        },
        "state_code": "CA"
    }
  }`
}

func makeFakeServer(resp string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, resp)
	}))
}
