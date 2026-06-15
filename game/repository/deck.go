package repository

import "godot-golang-server/model"

func NewDeck() *model.Deck {
	return &model.Deck{
		Deck: map[string][]model.Card{
			"teste1": {
				{Id: 1, Name: "Attack", Cost: 1, Text: "Deal 3 damage to opponent"},
				{Id: 2, Name: "Defense", Cost: 1, Text: "Heals 3 damage to you"},
				{Id: 1, Name: "Attack", Cost: 1, Text: "Deal 3 damage to opponent"},
				{Id: 1, Name: "Attack", Cost: 1, Text: "Deal 3 damage to opponent"},
			},
			"teste2": {
				{Id: 2, Name: "Defense", Cost: 1, Text: "Heals 3 damage to you"},
				{Id: 2, Name: "Defense", Cost: 1, Text: "Heals 3 damage to you"},
				{Id: 2, Name: "Defense", Cost: 1, Text: "Heals 3 damage to you"},
				{Id: 2, Name: "Defense", Cost: 1, Text: "Heals 3 damage to you"},
			},
		},
	}
}
