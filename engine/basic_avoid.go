package engine

import (
	"battlesnake-go-crazy/models"
	"log"
)

// BasicAvoid implements a simple game engine that does random things.
// It should die really fast ...
type BasicAvoid struct{}

func (b *BasicAvoid) Description() string {
	return "Basic Avoid"
}

func (b *BasicAvoid) Info() models.BattlesnakeInfoResponse {
	log.Println("INFO")
	return models.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "Nuno Lopes",
		Color:      "#888888",
		Head:       "default",
		Tail:       "default",
	}
}

func (b *BasicAvoid) Start(state models.GameState) {
	log.Printf("START %s\n", state.Game.ID)
}

func (b *BasicAvoid) Move(state models.GameState) models.BattlesnakeMoveResponse {
	moves := models.NewValidMoves()

	state.NoCollisions(state.You, moves)

	return models.BattlesnakeMoveResponse{Move: moves.Random(), Shout: ""}
}

func (b *BasicAvoid) End(state models.GameState) {
	log.Printf("END %s\n----------\n", state.Game.ID)
}
