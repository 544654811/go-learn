package tcpServer

import (
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read data failed, err:", err)
		}

		data := string(buf[:n])
		fmt.Printf("accept cline send data, data: %s \n", data)
	}
}

func Test() {
	server, err := net.Listen("tcp", "0.0.0.0:20000")
	if err != nil {
		fmt.Println("tcpServer connet failed, err:", err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("tcpServer accept connet failed, err:", err)
			continue
		}

		go handleConn(conn)
	}
}
