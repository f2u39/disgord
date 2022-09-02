package server

import (
	"fmt"
	"log"
	"net"

	"disgord/user"
)

type Server struct {
	ip   string
	port int

	// A channel for holds incoming msgChan
	msgChan chan []byte

	// A map holds all users in the server
	users map[*user.User]bool

	// A channel for user inChan the server
	inChan chan *user.User

	// A channel for user outChan the server
	outChan chan *user.User
}

func NewServer(ip string, port int) *Server {
	return &Server{
		ip:      ip,
		port:    port,
		msgChan: make(chan []byte),
		users:   make(map[*user.User]bool),
		inChan:  make(chan *user.User),
		outChan: make(chan *user.User),
	}
}

func (s *Server) Select() {
	for {
		select {
		case user := <-s.inChan:
			s.users[user] = true
			fmt.Println(s.users)
		case user := <-s.outChan:
			s.users[user] = false
			fmt.Println(s.users)
		case msg := <-s.msgChan:
			for user := range s.users {
				user.MsgChan <- msg
			}
		}
	}
}

func (s *Server) ReceiveMsg(conn net.Conn) {
	for {
		msg := make([]byte, 1024)
		_, err := conn.Read(msg)
		if err != nil {
			log.Println(err)
		}
		println("Server receives message:", string(msg))

		s.msgChan <- msg
	}
}

func (s *Server) Serve() {
	lis, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		fmt.Println(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go s.ReceiveMsg(conn)

		u := user.NewUser(conn)
		s.inChan <- u
	}
}
