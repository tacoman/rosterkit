package models

type Player struct {
	Name	string `json:"name"`
	SquadNumber	string `json:"squadNumber"`
	Position	string `json:"position"`
}
type Foe struct {
	Opponent	string `json:"opponent"`
	Players		[]Player `json:"players"`
}

type FoeDef struct {
	Url string
	Name string
}