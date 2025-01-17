package main

import (
	"fmt"
	"sync"
)

// Node represents a node in the mesh topology.
type Node struct {
	id      int
	channel chan string
}

// Mesh represents the network in the mesh topology.
type Mesh struct {
	nodes []*Node
	wg    *sync.WaitGroup
	done  chan struct{} // Signal for stopping goroutines
}

// Create a new Node.
func newNode(id int) *Node {
	return &Node{
		id:      id,
		channel: make(chan string),
	}
}

// Create a new Mesh.
func newMesh() *Mesh {
	return &Mesh{
		nodes: []*Node{},
		wg:    &sync.WaitGroup{},
		done:  make(chan struct{}), // Signal channel
	}
}

// Add a node to the mesh.
func (mesh *Mesh) addNode(node *Node) {
	mesh.nodes = append(mesh.nodes, node)
}

// Connect all nodes in the mesh.
func (mesh *Mesh) connectNodes() {
	for _, sender := range mesh.nodes {
		mesh.wg.Add(1)
		go func(sender *Node) {
			defer mesh.wg.Done()
			for {
				select {
				case msg, ok := <-sender.channel:
					if !ok {
						// Channel closed, stop processing
						return
					}
					fmt.Printf("Node %d sent: %s\n", sender.id, msg)
					for _, receiver := range mesh.nodes {
						if receiver.id != sender.id {
							go func(receiver *Node, msg string) {
								receiver.channel <- fmt.Sprintf("From Node %d: %s", sender.id, msg)
							}(receiver, msg)
						}
					}
				case <-mesh.done:
					// Stop signal received
					return
				}
			}
		}(sender)
	}
}

// Start each node to listen for incoming messages.
func (node *Node) startListening(mesh *Mesh) {
	mesh.wg.Add(1)
	go func() {
		defer mesh.wg.Done()
		for {
			select {
			case msg, ok := <-node.channel:
				if !ok {
					// Channel closed, stop processing
					return
				}
				fmt.Printf("Node %d received: %s\n", node.id, msg)
			case <-mesh.done:
				// Stop signal received
				return
			}
		}
	}()
}

func main() {
	mesh := newMesh()

	// Create and add nodes to the mesh.
	for i := 1; i <= 3; i++ { // Adjust the number of nodes here.
		node := newNode(i)
		mesh.addNode(node)
		node.startListening(mesh)
	}

	// Connect all nodes in the mesh.
	mesh.connectNodes()

	// Simulate message sending.
	go func() {
		for i := 1; i <= 4; i++ {
			for _, node := range mesh.nodes {
				node.channel <- fmt.Sprintf("Message %d from Node %d", i, node.id)
			}
		}
		// Signal completion of message sending.
		close(mesh.done)
	}()

	// Wait for all communication to complete.
	mesh.wg.Wait()
	fmt.Println("Mesh topology simulation completed.")
}
