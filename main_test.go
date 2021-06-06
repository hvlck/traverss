package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mmcdole/gofeed"
)

func TestJson(t *testing.T) {
	resp, _ := http.Get("/json/drewdevault.com/blog/index.xml")

	var x gofeed.Feed
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&x)

	if x.Title != "Drew DeVault's blog" {
		t.Fail()
	}
}
