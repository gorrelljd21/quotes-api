package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type quote struct {
	ID     string `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

type ID struct {
	ID string `json:"id"`
}

var db *sql.DB

func main() {
	err := databaseConnection()
	if err != nil {
		log.Println(err)
	}

	r := gin.Default()
	// r.GET("/quotes", getRandomQuote)
	r.GET("/quotes", getRandomQuoteSQL)
	// r.GET("/quotes/:id", getQuoteById)
	r.GET("/quotes/:id", getQuoteByIdSQL)
	// r.POST("/quotes", addQuote)
	r.POST("/quotes", addQuoteSQL)
	r.Run("0.0.0.0:8080")
}

func databaseConnection() error {
	mustGetenv := func(dns string) string {
		gettingEnv := os.Getenv(dns)
		if gettingEnv == "" {
			log.Printf("Warning: %s environment variable not set", dns)
		}
		return gettingEnv
	}

	var (
		dbUser         = os.Getenv("DB_USER") //postgres
		dbPwd          = mustGetenv("DB_PWD")
		dbName         = mustGetenv("DB_NAME")              //postgres
		unixSocketPath = mustGetenv("INSTANCE_UNIX_SOCKET") // /cloudsql/jessie-apprentice:us-central1:quotes-database
	)

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s", dbUser, dbPwd, dbName, unixSocketPath)

	//dbPool is the pool of database connections
	var err error

	db, err = sql.Open("pgx", dbURI)
	if err != nil {
		return fmt.Errorf("sql.Open: %v", err)
	}
	return err
}

func manageHeader(c *gin.Context) bool {
	headers := c.Request.Header
	header, exists := headers["X-Api-Key"]
	fmt.Println(header)

	if exists {
		if header[0] == "COCKTAILSAUCE" {
			return true
		}
	}
	return false
}

func getRandomQuoteSQL(c *gin.Context) {
	row := db.QueryRow("select id, phrase, author from quotes order by RANDOM() limit 1")
	q := &quote{}
	err := row.Scan(&q.ID, &q.Quote, &q.Author)

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, q)
	return
}

// TODO fix invalid id err handler
func getQuoteByIdSQL(c *gin.Context) {
	if manageHeader(c) {
		id := c.Param("id")
		row, exists := db.QueryRow(fmt.Sprintf("select id, phrase, author from quotes where id = '%s'", id)), true
		q := &quote{}
		err := row.Scan(&q.ID, &q.Quote, &q.Author)

		if err != nil {
			log.Println(err)
		}

		if exists {
			c.JSON(http.StatusOK, q)
			return
		} else if !exists {
			c.JSON(http.StatusNotFound, gin.H{"message": "invalid ID"})
		}
	} else if !manageHeader(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
	}
}

// TODO authentication
func addQuoteSQL(c *gin.Context) {
	var newID ID
	q := &quote{}

	newID.ID = uuid.New().String()

	if flaw := c.BindJSON(&q); flaw != nil {
		return
	}

	sqlStatement := `insert into quotes (id, phrase, author) values ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, &newID.ID, &q.Quote, &q.Author)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusCreated, newID)
}

// func addQuote(c *gin.Context) {

// 	if manageHeader(c) {
// 		var newQuote quote
// 		var newID ID

// 		if err := c.BindJSON(&newQuote); err != nil {
// 			return
// 		}

// 		newUUID := uuid.New()
// 		newQuote.ID = newUUID.String()
// 		newID.ID = newUUID.String()

// 		if len(newQuote.Quote) < 3 || len(newQuote.Author) < 3 {
// 			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
// 			return
// 		} else {
// 			mapOfQuotes[newQuote.ID] = newQuote
// 			c.JSON(http.StatusCreated, newID)

// 		}
// 	} else if !manageHeader(c) {
// 		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
// 	}
// }

var mapOfQuotes = map[string]quote{
	"0d949b68-6b04-4b35-82e5-63159b7608f8": {ID: "0d949b68-6b04-4b35-82e5-63159b7608f8", Quote: "Don't communicate by sharing memory, share memory by communicating.", Author: "Rob Pike"},
	"f9a7ab3a-9fc5-40b3-8c2e-76239ca037ce": {ID: "f9a7ab3a-9fc5-40b3-8c2e-76239ca037ce", Quote: "Concurrency is not parallelism.", Author: "Rob Pike"},
	// "a240c7e9-1570-4c36-ae5f-699e4cb5e4d7": {ID: "a240c7e9-1570-4c36-ae5f-699e4cb5e4d7", Quote: "The bigger the interface, the weaker the abstraction.", Author: "Rob Pike"},
	"a2523b46-42d4-42f6-aeb9-42da4b928c4a": {ID: "a2523b46-42d4-42f6-aeb9-42da4b928c4a", Quote: "Use consistent spelling of certain words.", Author: "Dmitri Shuralyov"},
	"f5a05e7f-1e71-462f-8036-9b7c8bfbed65": {ID: "f5a05e7f-1e71-462f-8036-9b7c8bfbed65", Quote: "Single spaces between spaces.", Author: "Dmitri Shuralyov"},
	"7dbde6f1-c411-40ca-af84-cc7fec7c06ec": {ID: "7dbde6f1-c411-40ca-af84-cc7fec7c06ec", Quote: "Avoid unused method receiver names.", Author: "Dmitri Shuralyov"},
	"170f9d56-369e-4088-a23d-5c8bc3e4a973": {ID: "170f9d56-369e-4088-a23d-5c8bc3e4a973", Quote: "Comments for humans always have a single space after the slashes.", Author: "Dmitri Shuralyov"},
}

// func getQuoteById(c *gin.Context) {

// 	if manageHeader(c) {
// 		id := c.Param("id")

// 		quote, exists := mapOfQuotes[id]

// 		if exists {
// 			c.JSON(http.StatusOK, quote)
// 			return
// 		}
// 		c.JSON(http.StatusNotFound, gin.H{"message": "quote not found"})
// 	} else if !manageHeader(c) {
// 		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
// 	}
// }

// func getRandomQuote(c *gin.Context) {
// 	quoteSlice := []string{}

// 	if manageHeader(c) {
// 		for k := range mapOfQuotes {
// 			quoteSlice = append(quoteSlice, k)
// 		}
// 		randNum := rand.Intn(len(quoteSlice))
// 		randKey := quoteSlice[randNum]
// 		randQuote := mapOfQuotes[randKey]
// 		c.JSON(http.StatusOK, randQuote)
// 	} else if !manageHeader(c) {
// 		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
// 	}
// }
