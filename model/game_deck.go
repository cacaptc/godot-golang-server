package model

import (
	"errors"
	"math/rand"
)

// GameDeck representa o baralho DURANTE uma partida.
// É separado do Deck template porque tem estado mutável:
// quais cartas já foram compradas, ordem embaralhada, etc.
type GameDeck struct {
	Name  string // nome do deck
	Queue []Card // fila de cartas
	Drawn []Card // cartas já compradas
}

func NewGameDeck(deckName string, cards []Card) *GameDeck {
	// Copia as cartas para não modificar o template original
	queue := make([]Card, len(cards))
	copy(queue, cards)

	// Embaralha usando Fisher-Yates
	rand.Shuffle(len(queue), func(i, j int) {
		queue[i], queue[j] = queue[j], queue[i]
	})

	return &GameDeck{
		Name:  deckName,
		Queue: queue,
		Drawn: []Card{},
	}
}

func (gd *GameDeck) DrawCard() (Card, error) {
	if len(gd.Queue) == 0 {
		return Card{}, errors.New("baralho vazio")
	}

	card := gd.Queue[0]
	gd.Queue = gd.Queue[1:]

	gd.Drawn = append(gd.Drawn, card)

	return card, nil
}

func (gd *GameDeck) RemainingCards() int {
	return len(gd.Queue)
}
