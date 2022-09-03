package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"disgord/user"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (s *Server) Join(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

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

func (s *Server) Serve() {
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

func (s *Server) Serve0() {
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
