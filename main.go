package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/mmcdole/gofeed"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Documentation is available at https://github.com/hvlck/traverss")
}

func Json(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	url := "https://" + strings.TrimPrefix(r.URL.Path, "/json/")

	resp, err := http.Get(url)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"status":"failed","error":"failed to fetch url"}`)
		return
	}

	ct := resp.Header.Get("content-type")

	if !strings.Contains(ct, "application/atom+xml") && !strings.Contains(ct, "text/xml") && !strings.Contains(ct, "application/feed+json") && !strings.Contains(ct, "application/rss+xml") && !strings.Contains(ct, "application/xml") {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"status":"failed","error":"not rss"}`)
		return
	}

	parse := gofeed.NewParser()
	feed, err := parse.Parse(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"status":"failed","error":"failed to parse feed"}`)
		return
	}

	fmt.Fprint(w, feed.String())
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/json/", Json)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
