package connect

import (
	"github.com/charmbracelet/log"
	"net"
	"panzer_server/utils"
	"sync"
)

const (
	bindAddress = "127.0.0.1:8080"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr().String()
	for {
		tcpContent := make([]byte, 1024)
		_, err2 := conn.Read(tcpContent)
		tcpContentString := utils.RStripBinaryZeros(tcpContent)
		if err2 != nil {
			log.Error("Failed to read message, closing:", "connection", remoteAddr)
			conn.Close()
			return
		}

		responseData := processRequest(remoteAddr, tcpContentString)

		log.Debug("Get", "message", tcpContentString)
		_, err := conn.Write(responseData)
		if err != nil {
			log.Error("Failed to send message, closing connection", "conn", conn)
			conn.Close()
			return
		}
	}
}

func Listen() {

	log.Info("Try to listen on ", "address", bindAddress)
	ln, err := net.Listen("tcp", bindAddress)
	log.Info("Successfully to listening")
	if err != nil {
		log.Error("Failed to listening")
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
