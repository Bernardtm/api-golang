package socket

import (
	"bernardtm/backend/internal/core/shareds"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type PriorityLevel struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type NotificationType struct {
	ReceivedUUID  string         `json:"receivedUUID"`
	UUID          *string        `json:"uuid"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	CreationDate  string         `json:"creationDate"`
	PriorityLevel *PriorityLevel `json:"priorityLevel"`
	OptionType    string         `json:"type"`
}

type WebSocketHub struct {
	clients map[string]*websocket.Conn
	mu      sync.Mutex
}

var hub = WebSocketHub{
	clients: make(map[string]*websocket.Conn),
}

type SocketController interface {
	WebSocketHandler(c *gin.Context)
}

type socketController struct {
}

func NewSocketController() *socketController {
	return &socketController{}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (f *socketController) WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, shareds.ErrorResponse{Message: "Failed to upgrade to websocket connection"})
		return
	}

	userUUID, _ := c.Get("ID")

	hub.mu.Lock()
	hub.clients[userUUID.(string)] = conn
	hub.mu.Unlock()

	log.Printf("user %s connected", userUUID.(string))

	defer func() {
		hub.mu.Lock()

		delete(hub.clients, userUUID.(string))

		hub.mu.Unlock()

		conn.Close()
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}

}

func SendNotification(userUUID string, message NotificationType) {
	hub.mu.Lock()
	conn, ok := hub.clients[userUUID]
	hub.mu.Unlock()

	if !ok {
		log.Printf("Usuário %s não está conectado", userUUID)
		return
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Printf("Erro ao serializar mensagem para usuário %s: %v", userUUID, err)
		return
	}

	if ok {
		err := conn.WriteMessage(websocket.TextMessage, jsonMessage)
		if err != nil {
			log.Printf("Failed to send message to user %s: %v", userUUID, err)
			conn.Close()
			hub.mu.Lock()
			delete(hub.clients, userUUID)
			hub.mu.Unlock()
		}
	}
}
