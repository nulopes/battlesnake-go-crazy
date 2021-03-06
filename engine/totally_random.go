package engine

import (
	"battlesnake-go-crazy/models"
	"log"
)

// TotallyRandomEngine implements a simple game engine that does random things.
// It should die really fast ...
type TotallyRandomEngine struct{}

func (b *TotallyRandomEngine) Description() string {
	return "Totally Random"
}

func (b *TotallyRandomEngine) Info() models.BattlesnakeInfoResponse {
	log.Println("INFO")
	return models.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "Nuno Lopes",
		Color:      "#888888",
		Head:       "default",
		Tail:       "default",
	}
}

func (b *TotallyRandomEngine) Start(state models.GameState) {
	log.Printf("START %s\n", state.Game.ID)
}

func (b *TotallyRandomEngine) Move(_ models.GameState) models.BattlesnakeMoveResponse {
	moves := models.NewValidMoves()

	return models.BattlesnakeMoveResponse{Move: moves.Random(), Shout: ""}
}

func (b *TotallyRandomEngine) End(state models.GameState) {
	log.Printf("END %s\n----------\n", state.Game.ID)
}
