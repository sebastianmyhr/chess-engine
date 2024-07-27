package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Empty = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

const (
	White = iota
	Black
)

type Piece struct {
	Type  int `json:"type"`
	Color int `json:"color"`
}

type Board [8][8]*Piece

func NewBoard() *Board {
	board := &Board{}
	// Initialize the pieces on the board
	for i := 0; i < 8; i++ {
		board[1][i] = &Piece{Pawn, White}
		board[6][i] = &Piece{Pawn, Black}
	}
	board[0][0], board[0][7] = &Piece{Rook, White}, &Piece{Rook, White}
	board[7][0], board[7][7] = &Piece{Rook, Black}, &Piece{Rook, Black}
	board[0][1], board[0][6] = &Piece{Knight, White}, &Piece{Knight, White}
	board[7][1], board[7][6] = &Piece{Knight, Black}, &Piece{Knight, Black}
	board[0][2], board[0][5] = &Piece{Bishop, White}, &Piece{Bishop, White}
	board[7][2], board[7][5] = &Piece{Bishop, Black}, &Piece{Bishop, Black}
	board[0][3] = &Piece{Queen, White}
	board[7][3] = &Piece{Queen, Black}
	board[0][4] = &Piece{King, White}
	board[7][4] = &Piece{King, Black}

	return board
}

var board = NewBoard()

func getBoard(c *gin.Context) {
	c.JSON(http.StatusOK, board)
}

func makeMove(c *gin.Context) {
	var move struct {
		FromX int `json:"fromX"`
		FromY int `json:"fromY"`
		ToX   int `json:"toX"`
		ToY   int `json:"toY"`
	}
	if err := c.BindJSON(&move); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValidMove(board, move.FromX, move.FromY, move.ToX, move.ToY) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid move"})
		return
	}

	board[move.ToY][move.ToX] = board[move.FromY][move.FromX]
	board[move.FromY][move.FromX] = nil

	c.JSON(http.StatusOK, board)
}

func isValidPawnMove(board *Board, fromX, fromY, toX, toY int) bool {
	piece := board[fromY][fromX]
	direction := 1
	if piece.Color == Black {
		direction = -1
	}

	// Regular move
	if fromX == toX && toY == fromY+direction && board[toY][toX] == nil {
		return true
	}

	// Initial double move
	if fromX == toX && ((piece.Color == White && fromY == 1 && toY == 3) ||
		(piece.Color == Black && fromY == 6 && toY == 4)) &&
		board[toY][toX] == nil && board[fromY+direction][fromX] == nil {
		return true
	}

	// Capture
	if abs(fromX-toX) == 1 && toY == fromY+direction && board[toY][toX] != nil &&
		board[toY][toX].Color != piece.Color {
		return true
	}

	return false
}

func isValidKnightMove(fromX, fromY, toX, toY int) bool {
	dx := abs(toX - fromX)
	dy := abs(toY - fromY)
	return (dx == 2 && dy == 1) || (dx == 1 && dy == 2)
}

func isValidBishopMove(board *Board, fromX, fromY, toX, toY int) bool {
	if abs(toX-fromX) != abs(toY-fromY) {
		return false
	}
	return isDiagonalPathClear(board, fromX, fromY, toX, toY)
}

func isValidRookMove(board *Board, fromX, fromY, toX, toY int) bool {
	if fromX != toX && fromY != toY {
		return false
	}
	return isStraightPathClear(board, fromX, fromY, toX, toY)
}

func isValidQueenMove(board *Board, fromX, fromY, toX, toY int) bool {
	return isValidBishopMove(board, fromX, fromY, toX, toY) ||
		isValidRookMove(board, fromX, fromY, toX, toY)
}

func isValidKingMove(fromX, fromY, toX, toY int) bool {
	dx := abs(toX - fromX)
	dy := abs(toY - fromY)
	return dx <= 1 && dy <= 1
}

func isDiagonalPathClear(board *Board, fromX, fromY, toX, toY int) bool {
	dx := sign(toX - fromX)
	dy := sign(toY - fromY)
	x, y := fromX+dx, fromY+dy
	for x != toX && y != toY {
		if board[y][x] != nil {
			return false
		}
		x += dx
		y += dy
	}
	return true
}

func isStraightPathClear(board *Board, fromX, fromY, toX, toY int) bool {
	if fromX == toX {
		dy := sign(toY - fromY)
		for y := fromY + dy; y != toY; y += dy {
			if board[y][fromX] != nil {
				return false
			}
		}
	} else {
		dx := sign(toX - fromX)
		for x := fromX + dx; x != toX; x += dx {
			if board[fromY][x] != nil {
				return false
			}
		}
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
	}
	if x > 0 {
		return 1
	}
	return 0
}

func isValidMove(board *Board, fromX, fromY, toX, toY int) bool {
	// Check if the move is within the board
	if fromX < 0 || fromX > 7 || fromY < 0 || fromY > 7 ||
		toX < 0 || toX > 7 || toY < 0 || toY > 7 {
		return false
	}

	piece := board[fromY][fromX]
	if piece == nil {
		return false
	}

	// Check if the destination is not occupied by a piece of the same color
	if board[toY][toX] != nil && board[toY][toX].Color == piece.Color {
		return false
	}

	// Implement piece-specific move validation
	switch piece.Type {
	case Pawn:
		return isValidPawnMove(board, fromX, fromY, toX, toY)
	case Knight:
		return isValidKnightMove(fromX, fromY, toX, toY)
	case Bishop:
		return isValidBishopMove(board, fromX, fromY, toX, toY)
	case Rook:
		return isValidRookMove(board, fromX, fromY, toX, toY)
	case Queen:
		return isValidQueenMove(board, fromX, fromY, toX, toY)
	case King:
		return isValidKingMove(fromX, fromY, toX, toY)
	}

	return false
}

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.StaticFile("/", "./static/index.html")

	r.GET("/board", getBoard)
	r.POST("/move", makeMove)
	r.Run(":8080")
}
