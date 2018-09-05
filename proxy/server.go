package proxy

import (
	"io"
	"log"
	"net"
	"net/http"
)

// Request the server to open port.
func ServerListener(addr string) net.Listener {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("unable to register tcp forward: ", err)
	}
	return listen
}

func GetObject(server net.Listener, object string) error {

	handler := func(resp http.ResponseWriter, req *http.Request) {
		res, err := http.Get(object)
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(resp, res.Body)
	}

	defer server.Close()
	return http.Serve(server, http.HandlerFunc(handler))
}
