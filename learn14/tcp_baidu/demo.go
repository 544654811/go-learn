package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("baidu connect failed, err:", err)
	}
	defer conn.Close()

	msg := "GET / Http/1.1\r\n"
	msg += "HOST: www.baidu.com\r\n"
	msg += "connection: clost\r\n"
	msg += "\r\n\r\n"

	_, err = io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("baidu connect failed, err:", err)
	}

	var data [1023]byte
	for {
		n, err := conn.Read(data[:])
		if err != nil || n == 0 {
			fmt.Println("baidu reda failed, err:", err)
			break
		}
		fmt.Printf(string(data[:n]))
	}
}
