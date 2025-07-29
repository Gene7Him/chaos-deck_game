// Full WebSocket server entrypoint with router
package main

import (
	"log"
	"net/http"
	"chaos_deck/backend/internal/websocket"
)

func main() {
	http.HandleFunc("/ws", websocket.HandleWS)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}