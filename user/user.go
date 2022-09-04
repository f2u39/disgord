package user

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type User struct {
	Conn    *websocket.Conn
	MsgChan chan []byte
}

func NewUser(conn *websocket.Conn) *User {
	u := &User{
		Conn:    conn,
		MsgChan: make(chan []byte),
	}

	go u.KeepReceivingMsg()

	return u
}

func (u *User) KeepReceivingMsg() {
	// defer u.Conn.Close()
	for msg := range u.MsgChan {
		fmt.Println("Received msg from server →", msg)
	}
}

func (u *User) sendMsg(msg string) {
	err := u.Conn.WriteMessage(
		websocket.TextMessage,
		[]byte(msg))

	if err != nil {
		log.Println(err)
	}
}

func (u *User) Send(c echo.Context) error {
	// conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	// if err != nil {
	// 	return err
	// }
	// defer conn.Close()

	// u := user.NewUser(conn)

	// s.inChan <- u

	// go s.KeepListeningThisUser(u)
	msg := c.Param("msg")
	u.sendMsg(msg)

	return nil
}
