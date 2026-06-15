package actions

import (
	"fmt"
	"godot-golang-server/model"
)

func DrawCard(gs *model.GameState) model.ServerMessage {
	card, err := gs.GameDeck.DrawCard()
	if err != nil {
		return model.ServerMessage{
			Type:    "deck_empty",
			Success: false,
			Message: "Seu baralho está vazio!",
		}
	}

	return model.ServerMessage{
		Type:      "card_drawn",
		Success:   true,
		Card:      &card,
		Remaining: gs.GameDeck.RemainingCards(),
		Message:   fmt.Sprintf("Você comprou '%s'! Restam %d cartas.", card.Name, gs.GameDeck.RemainingCards()),
	}
}
