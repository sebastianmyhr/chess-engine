package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastianmyhr/chess-engine/pkg/chess"
	"net/http"
)

func main() 
{
	r := gin.Default()
	board := chess.NewBoard()

	r.GET("/", func(c *gin.Context)
	{
		c.HTML(http.StatusOK, "index.html", gin.H
		{
			"board": board,
		})
	})

	r.POST("/move", func(c *gin.Context)
	{
		var move struct
		{
			FromX int 'json:"fromX"'
			FromY int `json:"fromY"`
            ToX   int `json:"toX"`
            ToY   int `json:"toY"`
		}
		if c.BindJSON(&move) == nil
		{
			success := board.MovePiece(move.FromX, move.FromY, move.ToX, move.ToY)
			c.JSON(http.StatusOK, gin.H{"success": success})
		}
	})

	r.LoadHTMLGlob("web/templates/*")
	r.Static("/static", "./web/static")
	r.Run(":8080")
}