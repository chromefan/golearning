package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	test(1,"sdf")
	conn, err := net.Dial("tcp", "104.225.159.223:39520")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go func() {
		// 从标准输入流中接收输入数据
		input := bufio.NewScanner(os.Stdin)
		fmt.Printf("Please type in something:\n")
		// 逐行扫描
		for input.Scan() {
			line := input.Text()
			// 输入bye时 结束
			if line == "bye" {
				break
			}
			text := fmt.Sprintf("%s : %s\n", time.Now().Format("15:35:39\n"),line)
			_, err := io.WriteString(conn, text)
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	mustCopy(os.Stdout,conn)
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
func test(src,key string)  {
	fmt.Println(src,key)
}