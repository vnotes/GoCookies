package tools

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"

	"golang.org/x/crypto/ssh"
)

type Endpoint struct {
	Host string
	Port int
}

func (endpoint *Endpoint) String() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}

func PrivateKeyFile(file string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(signer)
}

type SSHTunnel struct {
	Local  *Endpoint
	Server *Endpoint
	Remote *Endpoint

	Config *ssh.ClientConfig
}

func (tunnel *SSHTunnel) Start(ch *chan bool) error {
	listener, err := net.Listen("tcp", tunnel.Local.String())
	*ch <- true
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go tunnel.forward(conn)

	}
}

func (tunnel *SSHTunnel) forward(localConn net.Conn) {
	serverConn, err := ssh.Dial("tcp", tunnel.Server.String(), tunnel.Config)
	if err != nil {
		return
	}

	remoteConn, err := serverConn.Dial("tcp", tunnel.Remote.String())
	if err != nil {
		return
	}
	copyConn := func(writer, reader net.Conn) {
		_, err := io.Copy(writer, reader)
		if err != nil {
			return
		}
	}

	go copyConn(localConn, remoteConn)
	go copyConn(remoteConn, localConn)
}
