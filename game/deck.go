package game

import (
	"fmt"
	"godot-golang-server/model"
)

func LoadGameDeck(catalog *model.Deck, deckName string) (*model.GameDeck, error) {
	cards, ok := catalog.Deck[deckName]
	if !ok {
		return nil, fmt.Errorf("deck '%s' não encontrado", deckName)
	}
	return model.NewGameDeck(deckName, cards), nil
}
