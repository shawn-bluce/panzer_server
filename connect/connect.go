package connect

import (
	"github.com/charmbracelet/log"
	"net"
	"sync"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	log.Info("Start handle connection", "conn", conn)
	for {
		tcpContent := make([]byte, 1024)
		log.Debug("Waiting for message", "conn", conn)
		_, err2 := conn.Read(tcpContent)
		if err2 != nil {
			log.Error("Failed to read message, closing connection", "conn", conn)
			conn.Close()
			return
		}
		log.Debug("Sending message", "conn", conn)
		_, err := conn.Write([]byte("Hello World! from server"))
		if err != nil {
			log.Error("Failed to send message, closing connection", "conn", conn)
			conn.Close()
			return
		}
	}
}

func Listen() {

	log.Info("Try to listen on port 127.0.0.1:8080")
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	log.Info("TCP socket listening on port 127.0.0.1:8080")
	if err != nil {
		log.Error("Failed to listen on port 127.0.0.1:8080")
		panic(err)
	}

	var wg sync.WaitGroup

	for {
		log.Info("Waiting for connection...")
		conn, err := ln.Accept()
		log.Info("Connection accepted!")
		if err != nil {
			continue // or return from loop if server is shutting down
		}

		wg.Add(1)

		go func() {
			handleClient(conn)
			wg.Done()
		}()
	}

	wg.Wait() // wait for connections to finish
}
