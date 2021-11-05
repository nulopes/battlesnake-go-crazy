package models

import (
	"log"
	"math/rand"
)

type GameState struct {
	Game  Game        `json:"game"`
	Turn  int         `json:"turn"`
	Board Board       `json:"board"`
	You   Battlesnake `json:"you"`
}

type Game struct {
	ID      string  `json:"id"`
	Ruleset Ruleset `json:"ruleset"`
	Timeout int32   `json:"timeout"`
}

type Ruleset struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Board struct {
	Height int           `json:"height"`
	Width  int           `json:"width"`
	Food   []Coord       `json:"food"`
	Snakes []Battlesnake `json:"snakes"`

	// Used in non-standard game modes
	Hazards []Coord `json:"hazards"`
}

type Battlesnake struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Health  int32   `json:"health"`
	Body    []Coord `json:"body"`
	Head    Coord   `json:"head"`
	Length  int32   `json:"length"`
	Latency string  `json:"latency"`

	// Used in non-standard game modes
	Shout string `json:"shout"`
	Squad string `json:"squad"`
}

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c Coord) Move(s string) Coord {
	switch s {
	case "up":
		return Coord{X: c.X, Y: c.Y + 1}
	case "down":
		return Coord{X: c.X, Y: c.Y - 1}
	case "right":
		return Coord{X: c.X + 1, Y: c.Y}
	case "left":
		return Coord{X: c.X - 1, Y: c.Y}
	}

	log.Fatal("invalid coord move string", s)
	return c
}

func (c Coord) Equals(other Coord) bool {
	return c.X == other.X && c.Y == other.Y
}

type BattlesnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

type BattlesnakeMoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout,omitempty"`
}

type ValidMoves struct {
	valid map[string]bool
}

func NewValidMoves() *ValidMoves {
	return &ValidMoves{valid: map[string]bool{"up": true, "right": true, "down": true, "left": true}}
}

func (v *ValidMoves) isValid(s string) bool {
	valid, ok := v.valid[s]
	if !ok {
		log.Fatal("missuse of valid", s)
	}
	return valid
}

func (v *ValidMoves) unset(s string) {
	if _, ok := v.valid[s]; !ok {
		log.Fatal("missuse of valid", s)
	}
	v.valid[s] = false
}

func (v *ValidMoves) ValidList() []string {
	possible := make([]string, 0, 4)

	for k, v := range v.valid {
		if v {
			possible = append(possible, k)
		}
	}
	return possible
}

func (v *ValidMoves) Random() string {
	possible := v.ValidList()

	if len(possible) == 0 {
		return "up"
	}
	return possible[rand.Intn(len(possible))]
}
