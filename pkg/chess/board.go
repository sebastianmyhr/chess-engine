package chess

type Color string
type PieceType string

const (
	White Color = "white"
	Black Color = "black"

	King   PieceType = "king"
	Queen  PieceType = "queen"
	Rook   PieceType = "rook"
	Bishop PieceType = "bishop"
	Knight PieceType = "knight"
	Pawn   PieceType = "pawn"
)

type Piece struct {
	Color Color
	Type  PieceType
}

type Board struct {
	Squares [8][8]*Piece
}

func NewBoard() *Board {
	board := &Board{}
	// Initialize the board with pieces
	board.initializePieces()
	return board
}

func (b *Board) initializePieces() {
	// Place white pieces
	b.Squares[0][0] = &Piece{Color: White, Type: Rook}
	b.Squares[0][1] = &Piece{Color: White, Type: Knight}
	b.Squares[0][2] = &Piece{Color: White, Type: Bishop}
	b.Squares[0][3] = &Piece{Color: White, Type: Queen}
	b.Squares[0][4] = &Piece{Color: White, Type: King}
	b.Squares[0][5] = &Piece{Color: White, Type: Bishop}
	b.Squares[0][6] = &Piece{Color: White, Type: Knight}
	b.Squares[0][7] = &Piece{Color: White, Type: Rook}
	for i := 0; i < 8; i++ {
		b.Squares[1][i] = &Piece{Color: White, Type: Pawn}
	}

	// Place black pieces
	b.Squares[7][0] = &Piece{Color: Black, Type: Rook}
	b.Squares[7][1] = &Piece{Color: Black, Type: Knight}
	b.Squares[7][2] = &Piece{Color: Black, Type: Bishop}
	b.Squares[7][3] = &Piece{Color: Black, Type: Queen}
	b.Squares[7][4] = &Piece{Color: Black, Type: King}
	b.Squares[7][5] = &Piece{Color: Black, Type: Bishop}
	b.Squares[7][6] = &Piece{Color: Black, Type: Knight}
	b.Squares[7][7] = &Piece{Color: Black, Type: Rook}
	for i := 0; i < 8; i++ {
		b.Squares[6][i] = &Piece{Color: Black, Type: Pawn}
	}
}
