package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}

	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages

	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)

	for input.Scan() {
		messager := make(chan string,1)
		messages <- who + ": " + input.Text()
		buffer := input.Text()
		log.Println(input.Text())
		go HeartBeating(conn,messager)
		go GravelChannel(buffer,messager)
	}
	// NOTE: ignoring potential errors from input.Err()
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}
func GravelChannel(data string, mess chan string) {
	mess <- data
	close(mess)
}
func HeartBeating(conn net.Conn,readerChannel chan string) {
	select {
	case fk := <- readerChannel:
		conn.SetReadDeadline(time.Now().Add(time.Duration(100) * time.Second))
		log.Println(conn.RemoteAddr().String(), "readerChannel:", string(fk))
		break
	case <- time.After(100 * time.Second):
		log.Println(conn.RemoteAddr().String(), "timeout 5s")
		conn.Close()
	}
}
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(listener.Addr().String())
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
