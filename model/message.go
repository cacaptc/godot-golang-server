package model

// Mensagem recebida do cliente Godot
type ClientMessage struct {
	Type   string `json:"type"`
	Action string `json:"action"`
}

// Mensagem enviada para o cliente Godot
type ServerMessage struct {
	Type      string `json:"type"`
	Success   bool   `json:"success"`
	Card      *Card  `json:"card,omitempty"` // ponteiro: null se não houver carta
	Remaining int    `json:"remaining"`

	Message string `json:"message,omitempty"`
}
