package proxy

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/user"

	gh "github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
)

func username() string {
	current, err := user.Current()
	if err != nil {
		log.Fatalf("unable to determine username: %v", err)
	}
	uname := current.Username
	return uname
}

func homedir() string {
	homedir, err := (gh.Dir())
	if err != nil {
		log.Fatalf("unable to get home dir: %v", err)
	}
	return homedir
}

func forward(localConn net.Conn, remoteConn *ssh.Client) {

	fmt.Println("Should forward recived bytes")
}

func SSHConnect() {
	var user, host, port, keypath, bind string
	flag.StringVar(&user, "user", username(), "an username to login")
	flag.StringVar(&host, "host", "", "remote host")
	flag.StringVar(&port, "port", "22", "remote port")
	flag.StringVar(&keypath, "key", homedir()+"/.ssh/id_rsa", "path to private key")
	flag.StringVar(&bind, "bind", "localhost:8080", "proxy bind address")

	flag.Parse()

	if host == "" {
		flag.PrintDefaults()
		fmt.Println("remote host is mandatory")
		os.Exit(1)
	}

	// A public key may be used to authenticate against the remote
	// server by using an unencrypted PEM-encoded private key file.
	//
	// If you have an encrypted private key, the crypto/x509 package
	// can be used to decrypt it.
	key, err := ioutil.ReadFile(keypath)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote server and perform the SSH handshake.
	connection := host + ":" + port
	client, err := ssh.Dial("tcp", connection, config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	} else {
		fmt.Printf("Successful connection to: %v@%v:%v\n", user, host, port)
	}

	localListener, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("Error: Failed to bind to %s (%s)\n", bind, err)
	} else {
		fmt.Printf("Listening on %s\n", bind)
	}

	for {
		// Setup localConn (type net.Conn)
		localConn, err := localListener.Accept()
		if err != nil {
			log.Fatalf("listen.Accept failed: %v", err)
		}
		go forward(localConn, client)
	}

	defer client.Close()

	fmt.Print("Connection closed")
}
