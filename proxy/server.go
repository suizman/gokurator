package proxy

import (
	"fmt"
	"io"
	"net/http"
)

type HealthCheckResponse struct {
	Version int    `json:"version"`
	Status  string `json:"status"`
}

// Request the server to open port.
func NewServer(addr string) error {
	return http.ListenAndServe(addr, nil)
}

func AuthWrapper(h http.HandlerFunc, ApiKey string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if Api-Key header is empty
		if r.Header.Get("Api-Key") == "" {
			http.Error(w, "Missing Api-Key header", http.StatusUnauthorized)
			return
		}

		if r.Header.Get("Api-Key") != ApiKey {
			http.Error(w, "Bad credentials", http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r) // call original
	})

}

func GetObject(w http.ResponseWriter, req *http.Request) {

	object := req.Header.Get("Object")

	if req.Header.Get("Object") == "" {
		http.Error(w, "Missing Object header", http.StatusUnauthorized)
		return
	}

	res, err := http.Get(object)
	if err != nil {
		fmt.Println(err)
	}

	io.Copy(w, res.Body)

}
