package main

import (
	"fmt"
	"sync"
)

// Node represents a node in the star topology.
type Node struct {
	id      int
	channel chan string
	nodes   []*Node
}

// Hub represents the central node in the star topology.
type Hub struct {
	nodes    []*Node
	messages chan string
	wg       *sync.WaitGroup
}

func newNode(id int) *Node {
	return &Node{
		id:      id,
		channel: make(chan string),
		nodes:   make([]*Node, 0),
	}
}

func newHub() *Hub {
	return &Hub{
		nodes:    []*Node{},
		messages: make(chan string),
		wg:       &sync.WaitGroup{},
	}
}

// Hub Method - Add new node to the hub
func (hub *Hub) addNode(node *Node) {
	hub.nodes = append(hub.nodes, node)
}

// Hub method - Start the hub
func (hub *Hub) start() {
	go func() {
		defer func() { //close all nodes channels when the hub stops
			for _, node := range hub.nodes {
				close(node.channel)
			}
		}()
		for msg := range hub.messages { //listens for messages on the messages channel
			fmt.Println("Hub received:", msg)
			for _, node := range hub.nodes { //broadcasts each message to all connected nodes by sending it to their channel
				node.channel <- msg
			}
		}
	}()
}

// Node method - listen for messages
func (node *Node) listen(hub *Hub) {
	hub.wg.Add(1)
	go func() {
		defer hub.wg.Done()             //decrement the waitgroup counter
		for msg := range node.channel { //listens for messages on its channel in a goroutine
			fmt.Printf("Node %d received: %s\n", node.id, msg)
		}
	}()
}

func main() {
	hub := newHub() //initialize the hub

	// Create and add nodes to the hub
	for i := 1; i <= 5; i++ {
		node := newNode(i)
		hub.addNode(node)
		node.listen(hub) //each node starts listening to for the messages
	}

	hub.start() //start the hub

	go func() {
		for i := 1; i <= 10; i++ { //send 10 messages to the hub in a separate goroutine
			hub.messages <- fmt.Sprintf("Message from myself %d", i)
		}
		close(hub.messages) //close messages channel when all messages are sent
	}()
	hub.wg.Wait()
	fmt.Println("Star topology simulation completed without deadlock.")
}
