package board

type Board struct {
	Data [8][8]int
}

func BuildBoard() [8][8]int {
	var board [8][8]int
	for v := 0; v < 8; v++ {
		if v%2 == 0 {
			board[v] = [8]int{0, 1, 0, 1, 0, 1, 0, 1}
		} else {
			board[v] = [8]int{1, 0, 1, 0, 1, 0, 1, 0}
		}
	}

	return board
}
