package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
	UUID   string `json:"UUID"`
}

func main() {
	r := gin.Default()
	r.GET("/quote", getRandomQuote)
	r.Run("localhost:8080")

	id := uuid.New()
	fmt.Println(id.String())
}

var mapOfQuotes = map[int]quote{
	0: {UUID: "0d949b68-6b04-4b35-82e5-63159b7608f8", Quote: "Don't communicate by sharing memory, share memory by communicating.", Author: "Rob Pike"},
	1: {UUID: "f9a7ab3a-9fc5-40b3-8c2e-76239ca037ce", Quote: "Concurrency is not parallelism.", Author: "Rob Pike"},
	2: {UUID: "1f9c3697-5232-45a8-82b7-ba9ac5f0799c", Quote: "Channels orchestrate; mutexes serialize.", Author: "Rob Pike"},
	3: {UUID: "a240c7e9-1570-4c36-ae5f-699e4cb5e4d7", Quote: "The bigger the interface, the weaker the abstraction.", Author: "Rob Pike"},
	4: {UUID: "a2523b46-42d4-42f6-aeb9-42da4b928c4a", Quote: "Use consistent spelling of certain words.", Author: "Dmitri Shuralyov"},
	5: {UUID: "f5a05e7f-1e71-462f-8036-9b7c8bfbed65", Quote: "Single spaces between spaces.", Author: "Dmitri Shuralyov"},
	6: {UUID: "7dbde6f1-c411-40ca-af84-cc7fec7c06ec", Quote: "Avoid unused method receiver names.", Author: "Dmitri Shuralyov"},
	7: {UUID: "170f9d56-369e-4088-a23d-5c8bc3e4a973", Quote: "Comments for humans always have a single space after the slashes.", Author: "Dmitri Shuralyov"},
}

// func getQuote(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, quotes)
// }

func getRandomQuote(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(len(mapOfQuotes))
	randQuote := mapOfQuotes[randNum]
	c.IndentedJSON(http.StatusOK, randQuote)
}

// var quotes = []quote{
// 	{Author: "Rob Pike", Quote: "Don't communicate by sharing memory, share memory by communicating."},
// 	{Author: "Rob Pike", Quote: "Concurrency is not parallelism."},
// 	{Author: "Rob Pike", Quote: "Channels orchestrate; mutexes serialize."},
// 	{Author: "Rob Pike", Quote: "The bigger the interface, the weaker the abstraction."},
// 	{Author: "Dmitri Shuralyov", Quote: "Use consistent spelling of certain words."},
// 	{Author: "Dmitri Shuralyov", Quote: "Single spaces between spaces."},
// 	{Author: "Dmitri Shuralyov", Quote: "Avoid unused method receiver names."},
// 	{Author: "Dmitri Shuralyov", Quote: "Comments for humans always have a single space after the slashes."},
// }
