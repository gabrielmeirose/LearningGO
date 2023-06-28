package main

import (
	b "chess/board"
	"fmt"
)

var game_board b.Board = b.Board{
	Data: b.BuildBoard(),
}

func main() {
	for x := 0; x < len(game_board.Data); x++ {
		fmt.Println(game_board.Data[x])
	}
}
