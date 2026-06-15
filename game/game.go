package game

import (
	"fmt"
	"godot-golang-server/game/actions"
	"godot-golang-server/model"
)

func NewGame(catalog *model.Deck, deckName string) (*model.GameState, error) {
	gameDeck, err := LoadGameDeck(catalog, deckName)
	if err != nil {
		return nil, fmt.Errorf("erro ao iniciar jogo: %w", err)
	}

	return &model.GameState{
		Status:   model.StatusRunning,
		GameDeck: gameDeck,
	}, nil
}

func ProcessAction(gs *model.GameState, msg model.ClientMessage) model.ServerMessage {
	if gs.Status != model.StatusRunning {
		return model.ServerMessage{
			Type:    "error",
			Success: false,
			Message: "Partida não está em andamento.",
		}
	}

	switch msg.Action {
	case "draw_card":
		return actions.DrawCard(gs)
	default:
		return model.ServerMessage{
			Type:    "error",
			Success: false,
			Message: "Ação desconhecida: " + msg.Action,
		}
	}
}
