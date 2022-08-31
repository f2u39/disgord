package server

import (
	client "disgord/client"
	"fmt"
	"log"
	"net"
)

type Server struct {
	ip   string
	port int

	// A channel for holds incoming msgChan
	msgChan chan []byte

	// A map holds all clients in the server
	clients map[*client.Client]bool

	// A channel for client checkin the server
	inChan chan *client.Client

	// A channel for client checkout the server
	outChan chan *client.Client
}

func NewServer(ip string, port int) *Server {
	return &Server{
		ip:      ip,
		port:    port,
		msgChan: make(chan []byte),
		clients: make(map[*client.Client]bool),
		inChan:  make(chan *client.Client),
		outChan: make(chan *client.Client),
	}
}

func (s *Server) Select() {
	for {
		select {
		case cli := <-s.inChan:
			s.clients[cli] = true
			fmt.Println(s.clients)
		case cli := <-s.outChan:
			s.clients[cli] = false
			fmt.Println(s.clients)
		case msg := <-s.msgChan:
			for cli := range s.clients {
				cli.MsgChan <- msg
			}
		}
	}
}

func (s *Server) ReceiveMsg() {

}

func (s *Server) Serve() {
	lis, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		fmt.Println(err)
	}
	defer lis.Close()

	go s.Select()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		u := client.NewClient(conn)
		s.inChan <- u

		if conn == nil {
			s.outChan <- u
		}
	}
}
