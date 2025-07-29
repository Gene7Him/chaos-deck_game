// Room creation, player joins, turn order
package room

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Player struct {
	Conn *websocket.Conn
	ID   string
}

type Room struct {
	Players []*Player
	Mutex   sync.Mutex
}

var rooms = map[string]*Room{}

func NewPlayer(conn *websocket.Conn) *Player {
	return &Player{Conn: conn, ID: generatePlayerID()}
}

func AssignToRoom(player *Player) string {
	roomID := "room1" // Simplified for MVP
	if rooms[roomID] == nil {
		rooms[roomID] = &Room{Players: []*Player{}}
	}
	rooms[roomID].Mutex.Lock()
	defer rooms[roomID].Mutex.Unlock()
	rooms[roomID].Players = append(rooms[roomID].Players, player)
	return roomID
}

func GiveCardToPlayer(player *Player, roomID string) {}
func Broadcast(roomID string, data []byte) {}
func BroadcastCardPlay(player *Player, card interface{}, roomID string) {}
func generatePlayerID() string { return "player123" }
