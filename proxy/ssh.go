package proxy

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"

	gh "github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
)

func Username() string {
	current, err := user.Current()
	if err != nil {
		log.Fatalf("unable to determine username: %v", err)
	}
	uname := current.Username
	return uname
}

func Homedir() string {
	homedir, err := (gh.Dir())
	if err != nil {
		log.Fatalf("unable to get home dir: %v", err)
	}
	return homedir
}

func SSHConnect(keypath, user, host, port string) {
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

	defer client.Close()
	fmt.Print("Connection closed")
}
