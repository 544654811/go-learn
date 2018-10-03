package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:20000")
	if err != nil {
		fmt.Println("tcpClient connect failed, err:", err)
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("tcpClient reader failed, err:", err)
			break
		}

		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("tcpClient Write failed, err:", err)
			break
		}
	}
}
