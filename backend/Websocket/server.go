package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	launcher "mtgolauncher/backend/Launcher"
	flog "mtgolauncher/backend/Logging"
)

type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Details string `json:"details"`
}

var downloadInstance launcher.Download

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
	ModID     int
	FileURL   string
	ModName   string
	ModAuthor string
	Conn      *websocket.Conn
}

// Message is a WebSocket message structure.
type Message struct {
	Type      string `json:"type"`
	ModID     int    `json:"modID"`
	ModName   string `json:"modName"`
	FileURL   string `json:"fileURL"`
	ModAuthor string `json:"modAuthor"`
}

// InitWebSocket initializes the WebSocketManager and sets up routes.
func (wsm *WebSocketManager) InitWebSocket() {
	downloadInstance = *launcher.NewDownload()
	downloadInstance.StartDownloading()
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", wsm.handleWebSocket)

	wsm.server = &http.Server{
		Addr:    ":42145",
		Handler: mux,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		//		allowedOrigin := "http://localhost:34115"
		//		origin := r.Header.Get("Origin")
		return true
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
	for request := range wsm.download {
		downloadInstance.EnqueueDownload(request.ModID, request.FileURL, request.Conn, request.ModName, request.ModAuthor)
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
				ModID:     message.ModID,
				FileURL:   message.FileURL,
				ModName:   message.ModName,
				ModAuthor: message.ModAuthor,
				Conn:      conn,
			}
			wsm.download <- request
		default:
			log.Println("Unknown message type:", message.Type)
		}
	}
}

func (wsm *WebSocketManager) ShutdownServer() {
	if wsm.server != nil {
		fmt.Println("Shutting down WebSocket server...")
		if err := wsm.server.Shutdown(nil); err != nil {
			log.Println("Error shutting down WebSocket server:", err)
		} else {
			fmt.Println("WebSocket server has been shut down.")
		}
	}
}
