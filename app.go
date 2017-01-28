package main

import (
	"net/http"
	"os"

	xj "github.com/basgys/goxml2json"
	"github.com/rs/cors"
)

func main() {
	handler := cors.Default().Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		url.Scheme = "http"
		url.Host = "jws.jalan.net"
		q := url.Query()
		q.Set("key", os.Getenv("JWS_API_KEY"))
		url.RawQuery = q.Encode()
		res, err := http.Get(url.String())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json, err := xj.Convert(res.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(json.Bytes())
	}))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, handler)
}
