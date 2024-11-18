package api

import "net/http"

// https://docs.joinmastodon.org/methods/instance/
// https://fedidb.org/crawler.html

func Instance(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Instance\n"))
}
