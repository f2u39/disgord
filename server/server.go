package server

import (
	"fmt"
	"log"
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

type Server struct {
	ip   string
	port int

	// A channel for holds incoming message
	msgChan chan []byte

	// A map holds all users in the server
	users map[*user.User]bool

	// A channel for user join the server
	inChan chan *user.User

	// A channel for user leave the server
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
			fmt.Println("Someone joined →", s.users)

		case user := <-s.outChan:
			s.users[user] = false
			fmt.Println("Someone left →", s.users)

		case msg := <-s.msgChan:
			for user, ok := range s.users {
				if ok {
					user.MsgChan <- msg
					fmt.Println("Server send msg to user →", string(msg))
				}
			}
		}
	}
}

func (s *Server) KeepListeningThisUser(u *user.User) {
	for {
		_, msg, err := u.Conn.ReadMessage()
		if err != nil {
			s.outChan <- u
			return
		}

		fmt.Println("Message from client:", string(msg))
		s.msgChan <- msg
	}

	// for {
	// 	msg := make([]byte, 1024)
	// 	_, err := conn.Read(msg)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	println("Server receives message:", string(msg))

	// 	s.msgChan <- msg
	// }
}

func (s *Server) Join(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	// defer conn.Close()

	u := user.NewUser(conn)
	// fmt.Println("NewUser:", u.Conn)

	s.inChan <- u

	// defer func() { s.outChan <- u }()

	go s.KeepListeningThisUser(u)

	return nil
}

func (s *Server) Send(c echo.Context) error {
	msg := c.Param("msg")
	s.msgChan <- []byte(msg)
	return nil
}

// func (s *Server) Serve0() {
// 	lis, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", s.ip, s.port))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer lis.Close()

// 	for {
// 		conn, err := lis.Accept()
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}

// 		go s.ReceiveMsg(conn)

// 		u := user.NewUser(conn)
// 		s.inChan <- u
// 	}
// }

func Reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
