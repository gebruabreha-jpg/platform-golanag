package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"golang.org/x/crypto/ssh"
)

const (
	EXEC = "EXEC"
	ROOT = "ROOT"
	CONF = "CONF"
)

func main() {
	config := &ssh.ServerConfig{
		NoClientAuth: false,
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			log.Printf("Auth attempt: user=%s", c.User())
			if c.User() == "user" && string(pass) == "password" {
				log.Println("Auth SUCCESS")
				return nil, nil
			}
			log.Println("Auth FAILED")
			return nil, fmt.Errorf("auth failed")
		},
	}

	// Generate ECDSA P-256 key — widely supported by all SSH clients
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal("Failed to generate key:", err)
	}
	signer, err := ssh.NewSignerFromKey(privateKey)
	if err != nil {
		log.Fatal("Failed to create signer:", err)
	}
	config.AddHostKey(signer)
	log.Printf("Host key type: %s", signer.PublicKey().Type())

	listener, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	log.Println("Mock SSH server listening on :2222")
	log.Println("Credentials: user / password")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		go handleConn(conn, config)
	}
}

func handleConn(nConn net.Conn, config *ssh.ServerConfig) {
	sshConn, chans, reqs, err := ssh.NewServerConn(nConn, config)
	if err != nil {
		log.Println("Handshake/auth failed:", err)
		return
	}
	defer sshConn.Close()
	log.Printf("Connected: %s (user: %s)", sshConn.RemoteAddr(), sshConn.User())

	go ssh.DiscardRequests(reqs)

	for newChan := range chans {
		if newChan.ChannelType() != "session" {
			newChan.Reject(ssh.UnknownChannelType, "unsupported")
			continue
		}
		ch, requests, err := newChan.Accept()
		if err != nil {
			continue
		}
		go handleSession(ch, requests)
	}
}

func handleSession(ch ssh.Channel, reqs <-chan *ssh.Request) {
	defer ch.Close()

	// Accept all session requests (shell, pty, env)
	go func() {
		for req := range reqs {
			switch req.Type {
			case "shell", "pty-req", "env":
				if req.WantReply {
					req.Reply(true, nil)
				}
			default:
				if req.WantReply {
					req.Reply(false, nil)
				}
			}
		}
	}()

	shell := EXEC
	prompt := func() string {
		switch shell {
		case ROOT:
			return "root# "
		case CONF:
			return "(config)# "
		default:
			return "user$ "
		}
	}

	// Small delay to let pty-req/shell requests be processed
	ch.Write([]byte(prompt()))

	lineBuf := ""
	buf := make([]byte, 1024)

	for {
		n, err := ch.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println("Read error:", err)
			}
			return
		}

		lineBuf += string(buf[:n])

		for {
			idx := strings.IndexAny(lineBuf, "\r\n")
			if idx == -1 {
				break
			}
			cmd := strings.TrimSpace(lineBuf[:idx])
			lineBuf = strings.TrimLeft(lineBuf[idx:], "\r\n")

			if cmd == "" {
				ch.Write([]byte("\r\n" + prompt()))
				continue
			}

			log.Printf("[%s] cmd: %s", shell, cmd)

			switch {
			case cmd == "sudo su" && (shell == EXEC || shell == "SHELL"):
				ch.Write([]byte("\r\nPassword: "))
				pwBuf := ""
				for {
					pn, perr := ch.Read(buf)
					if perr != nil {
						return
					}
					pwBuf += string(buf[:pn])
					if strings.ContainsAny(pwBuf, "\r\n") {
						break
					}
				}
				pw := strings.TrimSpace(pwBuf)
				if pw == "password" {
					shell = ROOT
					ch.Write([]byte("\r\n" + prompt()))
				} else {
					ch.Write([]byte("\r\nAuth failure\r\n" + prompt()))
				}

			case cmd == "sh" && shell == EXEC:
				shell = "SHELL"
				ch.Write([]byte("\r\nuser$ "))

			case cmd == "configure" && shell == ROOT:
				shell = CONF
				ch.Write([]byte("\r\nEntering config mode\r\n" + prompt()))

			case cmd == "exit":
				switch shell {
				case CONF:
					shell = ROOT
				case ROOT:
					shell = EXEC
				case "SHELL":
					shell = EXEC
				default:
					ch.Write([]byte("\r\nlogout\r\n"))
					return
				}
				ch.Write([]byte("\r\n" + prompt()))

			case cmd == "whoami":
				user := "user"
				if shell == ROOT || shell == CONF {
					user = "root"
				}
				ch.Write([]byte("\r\n" + user + "\r\n" + prompt()))

			default:
				ch.Write([]byte("\r\n" + cmd + ": executed\r\n" + prompt()))
			}
		}
	}
}
