package user

import "net"

type User struct {
	Addr    string
	Conn    net.Conn
	MsgChan chan []byte
}

func NewUser(conn net.Conn) *User {
	addr := conn.RemoteAddr().String()

	u := &User{
		Conn:    conn,
		Addr:    addr,
		MsgChan: make(chan []byte),
	}

	go u.ReceiveMsg()

	return u
}

func (u *User) ReceiveMsg() {
	for {
		msg := <-u.MsgChan
		u.Conn.Write([]byte(msg))
	}
}
