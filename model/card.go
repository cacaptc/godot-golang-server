package model

type Card struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Cost int    `json:"cost"`
	Text string `json:"text"`
	Zone string `json:"Zone"`
}
