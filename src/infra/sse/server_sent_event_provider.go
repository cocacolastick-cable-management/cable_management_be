package sse

import (
	"bufio"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sync"
)

type ServerSentEventProvider struct {
	clients        map[string]chan string
	messageChannel chan string
	clientsLock    sync.Mutex
}

func NewSSEServer() *ServerSentEventProvider {
	return &ServerSentEventProvider{
		messageChannel: make(chan string, 100),
	}
}

func (s *ServerSentEventProvider) SendMessage(message string) {
	//s.clientsLock.Lock()
	//defer s.clientsLock.Unlock()
	//
	//// Find the channel for the client
	//channel, ok := s.clients[clientId]
	//if !ok {
	//	fmt.Printf("Client not found: %s\n", clientId)
	//	return
	//}

	// Send the message to the client's channel
	s.messageChannel <- message
}

//func (s *ServerSentEventProvider) HandleSSE(c *fiber.Ctx) error {
//	clientId := c.Query("client_id")
//
//	// Create a new channel for the client
//	messageChannel := make(chan string)
//	s.clientsLock.Lock()
//	s.clients[clientId] = messageChannel
//	s.clientsLock.Unlock()
//
//	// Set the response headers for Server-Sent Events
//	c.Set("Content-Type", "text/event-stream")
//	c.Set("Cache-Control", "no-cache")
//	c.Set("Connection", "keep-alive")
//	c.Set("Access-Control-Allow-Origin", "*")
//
//	// Send a welcome message to the client
//	fmt.Fprintf(c, "data: Welcome, client %s!\n\n", clientId)
//	c.Send(nil) // Manually send the response
//
//	// Wait for messages to send to the client
//	for {
//		message := <-messageChannel
//
//		// Send the message to the client
//		fmt.Fprintf(c, "data: %s\n\n", message)
//		c.Send(nil) // Manually send the response
//	}
//}

func (s *ServerSentEventProvider) HandleSSE(c *fiber.Ctx) error {
	// Set the response headers for Server-Sent Events
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	// Send a welcome message to the client
	fmt.Fprintf(c, "data: Welcome, client %s!\n\n", "vu")

	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		for {
			select {
			case message := <-s.messageChannel:
				fmt.Print("WRITTER")

				// Send the message to the client
				fmt.Fprintf(c, "data: %s\n\n", message)
				// Manually flush the response to send the message immediately
				// Flush the response to send the message immediately
				err := w.Flush()
				if err != nil {
					// Refreshing page in web browser will establish a new
					// SSE connection, but only (the last) one is alive, so
					// dead connections must be closed here.
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)

					break
				}
			}
		}
	})
	// Wait for messages to send to the client
	return nil
}
