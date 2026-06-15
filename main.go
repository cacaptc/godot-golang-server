package main

import (
	"godot-golang-server/handler"
	"log"
	"net/http"
)

func main() {
	// Registra a rota WebSocket
	http.HandleFunc("/ws", handler.HandleWebSocket)

	port := ":8080"
	log.Println("🚀 Servidor rodando em ws://localhost" + port + "/ws")

	// Inicia o servidor HTTP (que também serve WebSocket)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
