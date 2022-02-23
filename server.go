package main

import (
	"fmt"
	"net"
)

type Server struct {
	IP   string
	Port int
}

//Create a server interface
func NewServer(ip string, port int) *Server {
	server := &Server{
		IP:   ip,
		Port: port,
	}

	return server
}

func (this *Server) Handler(conn net.Conn) {
	//...currently connection business
	fmt.Println("Connection established successfully")
}

//The interface to start the server
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//close listen socket
	defer listener.Close()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		//do handler
		go this.Handler(conn)
	}
}
