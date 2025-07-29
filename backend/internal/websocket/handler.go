// WebSocket handlers for messages
package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"chaos_deck/backend/internal/game"
	"chaos_deck/backend/internal/room"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	player := room.NewPlayer(conn)
	roomID := room.AssignToRoom(player)

	gameLoop(player, roomID)
}

func gameLoop(player *room.Player, roomID string) {
	for {
		_, msg, err := player.Conn.ReadMessage()
		if err != nil {
			log.Println("ReadMessage error:", err)
			return
		}
		var message game.Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}
		game.ProcessMessage(message, player, roomID)
	}
}
