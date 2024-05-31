package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Manager struct {
	clients ClientList
	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		clients: make(ClientList),
	}
}

func (manager *Manager) wsHandler(res http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := NewClient(conn, manager)
	manager.addClient(client)
	go client.readMessages()
	go client.writeMessages()
}

func (manager *Manager) addClient(client *Client) {
	manager.Lock()
	defer manager.Unlock()
	manager.clients[client] = true
	fmt.Println("New client added!!")
}

func (manager *Manager) removeClient(client *Client) {
	manager.Lock()
	defer manager.Unlock()
	exists := manager.clients[client]
	if exists {
		client.connection.Close()
		delete(manager.clients, client)
	}
	fmt.Println("client removed!!")

}
