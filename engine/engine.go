package engine

import "battlesnake-go-crazy/models"

type Engine interface {
	Info() models.BattlesnakeInfoResponse
	Start(state models.GameState)
	Move(state models.GameState) models.BattlesnakeMoveResponse
	End(state models.GameState)
}
