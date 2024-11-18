package wellknown

import (
	"net/http"
)

// https://github.com/jhass/nodeinfo/tree/main
// https://fedidb.org/crawler.html

func NodeInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Node Information\n"))
}
