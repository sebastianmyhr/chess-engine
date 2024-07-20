package chess

func (b *Board) MovePiece(fromX, fromY, toX, toY int) bool {
	piece := b.Squares[fromX][fromY]
	if piece == nil {
		return false
	}

	// Implement move validation logic here
	if isValidMove(piece, fromX, fromY, toX, toY, b) {
		b.Squares[toX][toY] = piece
		b.Squares[fromX][fromY] = nil
		return true
	}

	return false
}

func isValidMove(piece *Piece, fromX, fromY, toX, toY int, board *Board) bool {
	// Implement specific movement rules for each piece type
	switch piece.Type {
	case King:
		return isValidKingMove(fromX, fromY, toX, toY)
	case Queen:
		return isValidQueenMove(fromX, fromY, toX, toY)
	case Rook:
		return isValidRookMove(fromX, fromY, toX, toY)
	case Bishop:
		return isValidBishopMove(fromX, fromY, toX, toY)
	case Knight:
		return isValidKnightMove(fromX, fromY, toX, toY)
	case Pawn:
		return isValidPawnMove(fromX, fromY, toX, toY, piece.Color, board)
	default:
		return false
	}
}

func isValidKingMove(fromX, fromY, toX, toY int) bool {
	// Kings move one square in any direction
	dx, dy := abs(toX-fromX), abs(toY-fromY)
	return dx <= 1 && dy <= 1
}

func isValidQueenMove(fromX, fromY, toX, toY int) bool {
	// Queens move like both a rook and a bishop
	return isValidRookMove(fromX, fromY, toX, toY) || isValidBishopMove(fromX, fromY, toX, toY)
}

func isValidRookMove(fromX, fromY, toX, toY int) bool {
	// Rooks move in straight lines
	return (fromX == toX || fromY == toY) && isPathClear(fromX, fromY, toX, toY)
}

func isValidBishopMove(fromX, fromY, toX, toY int) bool {
	// Bishops move diagonally
	dx, dy := abs(toX-fromX), abs(toY-fromY)
	return dx == dy && isPathClear(fromX, fromY, toX, toY)
}

func isValidKnightMove(fromX, fromY, toX, toY int) bool {
	// Knights move in an L-shape
	dx, dy := abs(toX-fromX), abs(toY-fromY)
	return (dx == 2 && dy == 1) || (dx == 1 && dy == 2)
}

func isValidPawnMove(fromX, fromY, toX, toY int, color Color, board *Board) bool {
	// Pawns move forward but capture diagonally
	direction := 1
	if color == Black {
		direction = -1
	}
	if toX == fromX+direction && fromY == toY && board.Squares[toX][toY] == nil {
		return true
	}
	if toX == fromX+direction && abs(toY-fromY) == 1 && board.Squares[toX][toY] != nil {
		return true
	}
	return false
}

func isPathClear(fromX, fromY, toX, toY int) bool {
	// Check if there are no pieces in the path
	dx := sign(toX - fromX)
	dy := sign(toY - fromY)
	x, y := fromX+dx, fromY+dy
	for x != toX || y != toY {
		if board.Squares[x][y] != nil {
			return false
		}
		x += dx
		y += dy
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}
