package models

// YouValidMoves returns the moves that do not collide with anything in the game
func (s GameState) NoCollisions(bs Battlesnake, vm *ValidMoves) {
	// avoid colliding with the walls
	s.avoidWalls(bs, vm)

	// avoid other Snakes
	for _, snake := range s.Board.Snakes {
		s.You.avoidBattlesnake(snake, vm)
	}
}

// avoidWalls avoids the virtual walls around the board
func (s GameState) avoidWalls(bs Battlesnake, vm *ValidMoves) {
	if vm.isValid("up") && bs.Head.Y >= s.Board.Height-1 {
		vm.unset("up")
	}
	if vm.isValid("down") && bs.Head.Y <= 0 {
		vm.unset("down")
	}

	if vm.isValid("right") && bs.Head.X >= s.Board.Width-1 {
		vm.unset("right")
	}
	if vm.isValid("left") && bs.Head.X <= 0 {
		vm.unset("left")
	}
}

// avoidBattlesnake avoids other battlesnakes
func (bs Battlesnake) avoidBattlesnake(bs2 Battlesnake, vm *ValidMoves) {
	possible := vm.ValidList()

	for _, move := range possible {
		mHead := bs.Head.Move(move)

		for _, bodyPart := range bs2.Body {
			if bodyPart.Equals(mHead) {
				vm.unset(move)
				break
			}
		}
	}
}
