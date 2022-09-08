package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func main() {
	r := gin.Default()
	r.GET("/quotes", getQuote)
	r.Run("localhost:8080")
}

func getQuote(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, quotes)
}

var quotes = []quote{
	{Author: "Rob Pike", Quote: "Don't communicate by sharing memory, share memory by communicating."},
	{Author: "Rob Pike", Quote: "Concurrency is not parallelism."},
	{Author: "Rob Pike", Quote: "Channels orchestrate; mutexes serialize."},
	{Author: "Rob Pike", Quote: "The bigger the interface, the weaker the abstraction."},
	{Author: "Dmitri Shuralyov", Quote: "Use consistent spelling of certain words."},
	{Author: "Dmitri Shuralyov", Quote: "Single spaces between spaces."},
	{Author: "Dmitri Shuralyov", Quote: "Avoid unused method receiver names."},
	{Author: "Dmitri Shuralyov", Quote: "Comments for humans always have a single space after the slashes."},
}
