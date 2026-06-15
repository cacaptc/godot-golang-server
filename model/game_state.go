package model

type GameState struct {
	Status   GameStatus
	GameDeck *GameDeck
	// Futuramente:
	// Player   *Player
	// Hand     []Card
	// Life     int
	// Mana     int
}

type GameStatus string

const (
	StatusWaiting  GameStatus = "waiting"   // aguardando início
	StatusRunning  GameStatus = "running"   // partida em andamento
	StatusGameOver GameStatus = "game_over" // partida encerrada
)
