// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
)

// !+broadcaster
type client struct {
	Channel chan<- string
	Name    string
	Address string
} // an outgoing message channel

func (c client) String() string {
	return c.Name + "(" + c.Address + ")"
}

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
				cli.Channel <- msg
			}

		case cli := <-entering:
			clients[cli] = true

			buff := bytes.Buffer{}
			buff.WriteString("Active clients:\n")
			for active, isActive := range clients {
				if isActive {
					buff.WriteString(active.String() + "\n")
				}
			}
			cli.Channel <- buff.String()

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Channel)
		}
	}
}

//!-broadcaster

// !+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	address := conn.RemoteAddr().String()
	ch <- "Your adress is " + address
	ch <- "What's your name:"

	input := bufio.NewScanner(conn)
	input.Scan()
	newClient := client{Channel: ch, Name: input.Text(), Address: address}

	ch <- "Welcome " + newClient.String()

	messages <- newClient.String() + " has arrived"
	entering <- newClient
	for input.Scan() {
		messages <- newClient.String() + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- newClient
	messages <- newClient.String() + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

// !+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

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
