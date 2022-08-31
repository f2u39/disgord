package client

import "net"

type Client struct {
	Addr    string
	Conn    net.Conn
	MsgChan chan []byte
}

func NewClient(conn net.Conn) *Client {
	addr := conn.RemoteAddr().String()
	return &Client{
		Conn:    conn,
		Addr:    addr,
		MsgChan: make(chan []byte),
	}
}

func (cli *Client) ReceiveMsg() {
	for {
		msg := <-cli.MsgChan
		cli.Conn.Write([]byte(msg))
	}
}
