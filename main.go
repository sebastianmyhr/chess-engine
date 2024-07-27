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

	board[move.ToY][move.ToX] = board[move.FromY][move.FromX]
	board[move.FromY][move.FromX] = nil

	c.JSON(http.StatusOK, board)
}

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.StaticFile("/", "./static/index.html")

	r.GET("/board", getBoard)
	r.POST("/move", makeMove)
	r.Run(":8080")
}
