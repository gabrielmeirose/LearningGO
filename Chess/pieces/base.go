package pieces

type Piece struct {
	Appearance string
	Color      bool
	Pos        [2]int
}

type Movement interface {
	Sight()
	Move()
}
