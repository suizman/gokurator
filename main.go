package main

import (
	"flag"

	p "github.com/suizman/goxyfy/proxy"
)

func main() {

	var user, host, port, keypath, bind, get string
	var server bool
	flag.BoolVar(&server, "server", false, "run in server mode")
	flag.StringVar(&user, "user", p.Username(), "an username to login")
	flag.StringVar(&host, "host", "", "remote host")
	flag.StringVar(&port, "port", "22", "remote port")
	flag.StringVar(&keypath, "key", p.Homedir()+"/.ssh/id_rsa", "path to private key")
	flag.StringVar(&bind, "bind", "localhost:8080", "proxy bind address")
	flag.StringVar(&get, "get", "http://example.com/", "the url from which you want to get the file")

	flag.Parse()

	if server == true {
		serve := p.ServerListener(bind)
		p.GetObject(serve, get)
	}

}
