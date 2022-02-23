package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

//Create a user API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	//Start a goroutine that listens for messages from the current user channel
	go user.ListenMessage()

	return user
}

//The method of monitoring the current User channel, it will be sent directly to the opposite client once there is a message.
//func (this *User) ListenMessage() {,
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
