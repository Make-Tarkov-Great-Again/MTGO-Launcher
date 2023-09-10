package websocket

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"

	launcher "mtgolauncher/backend/Launcher"
	flog "mtgolauncher/backend/Logging"
)

// WebSocketManager handles WebSocket connections and download requests.
type WebSocketManager struct {
	server   *http.Server
	download chan *DownloadRequest
}

// upgrader is a configuration for upgrading HTTP connections to WebSocket.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// DownloadRequest represents a request to download a file via WebSocket.
type DownloadRequest struct {
	ModID   int
	FileURL string
	Conn    *websocket.Conn
}

// Message is a WebSocket message structure.
type Message struct {
	Type    string `json:"type"`
	ModID   int    `json:"modID"`
	FileURL string `json:"fileURL"`
}

// InitWebSocket initializes the WebSocketManager and sets up routes.
func (wsm *WebSocketManager) InitWebSocket() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", wsm.handleWebSocket)

	wsm.server = &http.Server{
		Addr:    ":42145",
		Handler: mux,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		allowedOrigin := "http://localhost:34115"
		origin := r.Header.Get("Origin")
		return origin == allowedOrigin || strings.HasSuffix(origin, ".localhost:34115")
	}

	wsm.download = make(chan *DownloadRequest) //should i make more of these for other functions???
	// TODO: what else do i need this for
	//probaly updates and mod lists maybe kekw

	// haha go download
	go wsm.handleDownloadRequests()
}

// StartServer starts the WebSocket server.
func (wsm *WebSocketManager) StartServer() {
	err := wsm.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("Error starting WebSocket server: ", err)
	}
	fmt.Println("WebSocket has started.. Maybe.")
	flog.OnlineLog("WebSocket has started.")
}

// handleDownloadRequests processes download requests by communicating with the launcher.
func (wsm *WebSocketManager) handleDownloadRequests() {
	app := launcher.NewLauncher()
	for request := range wsm.download {
		err := app.Download.Mod(request.ModID, request.FileURL, request.Conn)
		if err != nil {
			log.Println("Error handling download:", err)
			app.UI.Error(fmt.Sprintf("Failed to download %s", request.FileURL), fmt.Sprintf("A error occured when downloading the following mod: \n %s \n\n %s", request.FileURL, err))
		}
	}
}

// handleWebSocket upgrades an HTTP connection to a WebSocket and handles incoming messages.
func (wsm *WebSocketManager) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	log.Println("Request Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			log.Printf("%s: %s\n", name, value)
		}
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		message := &Message{}
		err := conn.ReadJSON(message)
		if err != nil {
			log.Println("Error reading WebSocket message:", err)
			break
		}

		switch message.Type {
		case "download":
			request := &DownloadRequest{
				ModID:   message.ModID,
				FileURL: message.FileURL,
				Conn:    conn,
			}
			wsm.download <- request
		default:
			log.Println("Unknown message type:", message.Type)
		}
	}
}
