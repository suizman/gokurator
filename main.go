package main

import (
	"flag"
	"net/http"
	"os"

	c "github.com/suizman/goxyfy/client"
	p "github.com/suizman/goxyfy/proxy"
)

func main() {

	var user, host, proxy, port, keypath, bind, get, output string
	var server, client bool
	flag.BoolVar(&server, "server", false, "run in server mode")
	flag.BoolVar(&client, "client", false, "run client")
	flag.StringVar(&user, "user", p.Username(), "an username to login")
	flag.StringVar(&host, "host", "localhost:8080", "remote host")
	flag.StringVar(&proxy, "proxy", "http://localhost:8080", "remote host")
	flag.StringVar(&port, "port", "22", "remote port")
	flag.StringVar(&keypath, "key", p.Homedir()+"/.ssh/id_rsa", "path to private key")
	flag.StringVar(&bind, "bind", "localhost:8080", "proxy bind address")
	flag.StringVar(&get, "get", "http://example.com/", "the url from which you want to get the file")
	flag.StringVar(&output, "output", "output", "output file name")

	flag.Parse()

	if server == true {
		http.HandleFunc("/get", p.GetObject)
		p.Server(bind)
	}

	if client == true {
		c.GetObject(get, output, proxy)
		os.Exit(0)
	}
	flag.PrintDefaults()

}
