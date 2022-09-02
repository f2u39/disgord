package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	ip   = "127.0.0.1"
	port = ":5002"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Disgord! Plz input your name:")
	fmt.Println("----------------------------------------")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)

	addr, err := net.ResolveTCPAddr("tcp", ip+port)
	if err != nil {
		log.Println("failed to resolce tcp address", err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println("failed to connect to server", err)
		return
	}

	sendMsg(conn, "!name="+name)

	for {
		if receiveMsg(conn) != nil {
			return
		}
	}
}

func sendMsg(conn *net.TCPConn, msg string) {
	_, err := conn.Write([]byte(msg))
	if err != nil {
		log.Println(err)
		return
	}
}

func receiveMsg(conn *net.TCPConn) error {
	msg := make([]byte, 1024)
	_, err := conn.Read(msg)
	if err != nil {
		log.Println(err)
		return err
	}
	println(string(msg))
	return nil
}
