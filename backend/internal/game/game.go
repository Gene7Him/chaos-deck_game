// Game engine: play, draw, turn resolution
package game

import (
	"encoding/json"
	"log"
	"chaos_deck/backend/internal/room"
)

type Message struct {
	Type string `json:"type"`
	Data json.RawMessage `json:"data"`
}

type Card struct {
	Color  string `json:"color"`
	Value  string `json:"value"`
	Effect string `json:"effect,omitempty"`
}

func ProcessMessage(msg interface{}, player *room.Player, roomID string) {
	switch m := msg.(type) {
	case Message:
		switch m.Type {
		case "play_card":
			var card Card
			if err := json.Unmarshal(m.Data, &card); err != nil {
				log.Println("PlayCard decode error:", err)
				return
			}
			ApplyCardEffect(card, player, roomID)
		case "draw_card":
			room.GiveCardToPlayer(player, roomID)
		case "chat":
			room.Broadcast(roomID, m.Data)
		}
	default:
		log.Println("Unknown message type")
	}
}

func ApplyCardEffect(card Card, player *room.Player, roomID string) {
	switch card.Value {
	case "skip":
		skipNextTurn(roomID)
	case "reverse":
		reverseTurnOrder(roomID)
	case "draw2":
		forceDraw(player, roomID, 2)
	case "wild":
		setWildColor(card, player, roomID)
	case "trap":
		triggerTrap(player, roomID)
	default:
		room.BroadcastCardPlay(player, card, roomID)
	}
}

func skipNextTurn(roomID string) {}
func reverseTurnOrder(roomID string) {}
func forceDraw(player *room.Player, roomID string, count int) {}
func setWildColor(card Card, player *room.Player, roomID string) {}
func triggerTrap(player *room.Player, roomID string) {}
