package main

import (
	"github.com/gorilla/websocket"
	"log"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	egress     chan []byte
}

func NewClient(connection *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: connection,
		manager:    manager,
		egress:     make(chan []byte),
	}
}

func (client *Client) readMessages() {
	defer func() {
		client.manager.removeClient(client)
	}()

	for {
		messageType, payload, err := client.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error reading message:", err)
			}
			log.Println("error reading message 1:", err)
			break
		}
		for wsClient := range client.manager.clients {
			wsClient.egress <- payload
		}

		log.Println("message type : ", messageType)
		log.Println("payload : ", string(payload))
	}
}

func (client *Client) writeMessages() {
	defer func() {
		client.manager.removeClient(client)
	}()

	for {
		select {
		case message, ok := <-client.egress:
			if !ok {
				err := client.connection.WriteMessage(websocket.CloseMessage, nil)
				if err != nil {
					log.Println("connection closed : ", err)
				}
				return
			}
			if err := client.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("failed to send message : ", err)
			}
			log.Println("message sent")
		}
	}
}
