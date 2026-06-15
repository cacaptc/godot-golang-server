package handler

import (
	"encoding/json"
	"godot-golang-server/game"
	"godot-golang-server/game/repository"
	"godot-golang-server/model"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// catalog é carregado uma vez e compartilhado entre conexões (read-only)
var catalog = repository.NewDeck()

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erro ao fazer upgrade:", err)
		return
	}
	defer conn.Close()
	gs, err := game.NewGame(catalog, "teste1")
	if err != nil {
		log.Println("Erro ao criar jogo:", err)
		// sendError(conn, "Não foi possível iniciar o jogo.")
		return
	}

	log.Println("✅ Cliente conectado:", conn.RemoteAddr())

	// 2. Loop de escuta — fica rodando enquanto o cliente estiver conectado
	for {
		// Lê a próxima mensagem do cliente
		_, rawMsg, err := conn.ReadMessage()
		if err != nil {
			log.Println("❌ Cliente desconectado:", err)
			break
		}

		log.Printf("📨 Mensagem recebida: %s\n", rawMsg)

		// 3. Desserializa o JSON recebido
		var clientMsg model.ClientMessage
		if err := json.Unmarshal(rawMsg, &clientMsg); err != nil {
			log.Println("Erro ao parsear JSON:", err)
			continue
		}

		// 4. Processa a ação e monta a resposta
		response := game.ProcessAction(gs, clientMsg)

		// 5. Serializa a resposta em JSON e envia de volta
		responseJSON, err := json.Marshal(response)
		if err != nil {
			log.Println("Erro ao serializar resposta:", err)
			continue
		}

		if err := conn.WriteMessage(websocket.TextMessage, responseJSON); err != nil {
			log.Println("Erro ao enviar resposta:", err)
			break
		}

		log.Printf("📤 Resposta enviada: %s\n", responseJSON)
	}
}
func processAction(msg model.ClientMessage) model.ServerMessage {
	switch msg.Action {
	case "draw_card":
		return model.ServerMessage{
			Type:    "card_drawn",
			Success: true,
			Card:    nil,
			Message: "Você comprou uma carta!",
		}
	default:
		return model.ServerMessage{
			Type:    "error",
			Success: false,
			Message: "Ação desconhecida: " + msg.Action,
		}
	}
}
