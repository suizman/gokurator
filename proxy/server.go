package proxy

import (
	"fmt"
	"io"
	"net/http"
)

// Request the server to open port.
func Server(addr string) error {
	return http.ListenAndServe(addr, nil)
}

func GetObject(w http.ResponseWriter, req *http.Request) {
	object := req.Header.Get("object")
	if object == "" {
		fmt.Println("Empty object header")
	}
	res, err := http.Get(object)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(w, res.Body)
}
