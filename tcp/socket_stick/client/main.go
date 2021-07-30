package main

import (
	"app/golearning/tcp/socket_stick/proto"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 200000; i++ {
		msg := fmt.Sprintf("%s %d","Hello, Hello. How are you?",time.Now().Unix())
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
		time.Sleep(1*time.Second)
	}
	time.Sleep(1000*time.Second)
}