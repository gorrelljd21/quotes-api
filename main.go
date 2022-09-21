package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type quote struct {
	ID     string `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func main() {
	r := gin.Default()
	r.GET("/quotes", getRandomQuote)
	r.GET("/quotes/:id", getQuoteById)
	r.POST("/newQuote", addQuote)
	r.Run("0.0.0.0:8080")
}

func getRandomQuote(c *gin.Context) {

	apiKeySlice := c.Request.Header["X-Api-Key"]
	apiKeyString := apiKeySlice[0]

	if apiKeyString == "COCKTAILSAUCE" {
		quoteSlice := []string{}

		for k := range mapOfQuotes {
			quoteSlice = append(quoteSlice, k)
		}

		randNum := rand.Intn(len(quoteSlice))
		randKey := quoteSlice[randNum]
		randQuote := mapOfQuotes[randKey]
		c.JSON(http.StatusOK, randQuote)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "401"})
	}
}

func getQuoteById(c *gin.Context) {

	apiKeySlice := c.Request.Header["X-Api-Key"]
	apiKeyString := apiKeySlice[0]

	if apiKeyString == "COCKTAILSAUCE" {
		id := c.Param("id")

		quote, exists := mapOfQuotes[id]

		if exists {
			c.JSON(http.StatusOK, quote)
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "quote not found"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "401"})
	}

}

func addQuote(c *gin.Context) {

	apiKeySlice := c.Request.Header["X-Api-Key"]
	apiKeyString := apiKeySlice[0]

	if apiKeyString == "COCKTAILSAUCE" {
		var newQuote quote

		if err := c.BindJSON(&newQuote); err != nil {
			return
		}

		newUUID := uuid.New()
		newQuote.ID = newUUID.String()

		if len(newQuote.Quote) < 3 || len(newQuote.Author) < 3 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		} else {
			mapOfQuotes[newQuote.ID] = newQuote
			c.JSON(http.StatusCreated, newQuote)
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "401"})
	}
}

var mapOfQuotes = map[string]quote{
	"0d949b68-6b04-4b35-82e5-63159b7608f8": {ID: "0d949b68-6b04-4b35-82e5-63159b7608f8", Quote: "Don't communicate by sharing memory, share memory by communicating.", Author: "Rob Pike"},
	"f9a7ab3a-9fc5-40b3-8c2e-76239ca037ce": {ID: "f9a7ab3a-9fc5-40b3-8c2e-76239ca037ce", Quote: "Concurrency is not parallelism.", Author: "Rob Pike"},
	"1f9c3697-5232-45a8-82b7-ba9ac5f0799c": {ID: "1f9c3697-5232-45a8-82b7-ba9ac5f0799c", Quote: "Channels orchestrate; mutexes serialize.", Author: "Rob Pike"},
	"a240c7e9-1570-4c36-ae5f-699e4cb5e4d7": {ID: "a240c7e9-1570-4c36-ae5f-699e4cb5e4d7", Quote: "The bigger the interface, the weaker the abstraction.", Author: "Rob Pike"},
	"a2523b46-42d4-42f6-aeb9-42da4b928c4a": {ID: "a2523b46-42d4-42f6-aeb9-42da4b928c4a", Quote: "Use consistent spelling of certain words.", Author: "Dmitri Shuralyov"},
	"f5a05e7f-1e71-462f-8036-9b7c8bfbed65": {ID: "f5a05e7f-1e71-462f-8036-9b7c8bfbed65", Quote: "Single spaces between spaces.", Author: "Dmitri Shuralyov"},
	"7dbde6f1-c411-40ca-af84-cc7fec7c06ec": {ID: "7dbde6f1-c411-40ca-af84-cc7fec7c06ec", Quote: "Avoid unused method receiver names.", Author: "Dmitri Shuralyov"},
	"170f9d56-369e-4088-a23d-5c8bc3e4a973": {ID: "170f9d56-369e-4088-a23d-5c8bc3e4a973", Quote: "Comments for humans always have a single space after the slashes.", Author: "Dmitri Shuralyov"},
}
