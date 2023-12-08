package websocket

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// PENWSMessage represents a WebSocket message for popups, errors, and notices.
type PENWSMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// PENWSManager handles PENWS WebSocket connections and messages.
type PENWSManager struct {
	server  *http.Server
	clients map[*websocket.Conn]struct{}
	mu      sync.Mutex
}

// InitPENWS initializes the PENWSManager.
func (pwm *PENWSManager) InitPENWS() {
	pwm.clients = make(map[*websocket.Conn]struct{})
	mux := http.NewServeMux()
	mux.HandleFunc("/penws", pwm.handlePENWS)

	pwm.server = &http.Server{
		Addr:    ":42146", // Use a different port for PENWS
		Handler: mux,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	// TODO: Initialize any other PENWS-specific settings

	// Start PENWS server
	go pwm.startPENWSServer()
}

// startPENWSServer starts the PENWS server.
func (pwm *PENWSManager) startPENWSServer() {
	err := pwm.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("Error starting PENWS server: ", err)
	}
	fmt.Println("PENWS has started.. Probally lol.")
}

// handlePENWS upgrades an HTTP connection to PENWS and handles incoming messages.
func (pwm *PENWSManager) handlePENWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Add the client to the clients map
	pwm.mu.Lock()
	pwm.clients[conn] = struct{}{}
	pwm.mu.Unlock()

	for {
		message := &PENWSMessage{}
		err := conn.ReadJSON(message)
		if err != nil {
			log.Println("Error reading PENWS message:", err)

			// Remove the client from the clients map upon error
			pwm.mu.Lock()
			delete(pwm.clients, conn)
			pwm.mu.Unlock()

			break
		}

		// Handle the PENWS message, e.g., show a popup in the frontend
		switch message.Type {
		case "error":
			// Handle error message
			// You can implement logic here to show an error popup in the frontend
		case "notice":
			// Handle notice message
			// You can implement logic here to show a notice popup in the frontend
		default:
			log.Println("Unknown PENWS message type:", message.Type)
		}
	}
}

// BroadcastPENWSMessage sends a PENWS message to all connected clients.
func (pwm *PENWSManager) BroadcastPENWSMessage(message *PENWSMessage) {
	pwm.mu.Lock()
	defer pwm.mu.Unlock()

	for conn := range pwm.clients {
		err := conn.WriteJSON(message)
		if err != nil {
			log.Println("Error writing PENWS message:", err)
		}
	}
}

// ShutdownPENWSServer shuts down the PENWS server.
func (pwm *PENWSManager) ShutdownPENWSServer() {
	if pwm.server != nil {
		fmt.Println("Shutting down PENWS server...")
		if err := pwm.server.Shutdown(nil); err != nil {
			log.Println("Error shutting down PENWS server:", err)
		} else {
			fmt.Println("PENWS server has been shut down.")
		}
	}
}
