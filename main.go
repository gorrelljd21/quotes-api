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

type quote struct { //the point of this struct is to give JSON an id to look for in order to connect w the server
	ID     string `json:"id"` // object name type and JSON endpoint
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

type ID struct { //the point of this struct is to output just the id when a POST is used for a new quote
	ID string `json:"id"`
}

var db *sql.DB // hoisted this so it could be used anywhere below it

func main() { //the point of this function is to run the http requests
	err := databaseConnection() //29-32 block will check to see if the connection is there or not
	if err != nil {
		log.Println(err)
	}

	r := gin.Default() //34-39 is scripting the endpoints and connecting the matching functions that do the work. it identifies what server locally it would run on (39)
	r.GET("/quotes", getRandomQuoteSQL)
	r.GET("/quotes/:id", getQuoteByIdSQL)
	r.POST("/quotes", addQuoteSQL)
	r.DELETE("/quotes/:id", deleteQuote)
	r.Run("0.0.0.0:8080")
}

func databaseConnection() error { //the purpose of this function is to connect this golang code to cloud sql so it can read the database (the sql file is inputted within cloud sql)
	mustGetenv := func(dns string) string { // we add this default func so that the env's below can be inputted and checked to see if they are set or not
		gettingEnv := os.Getenv(dns) // this get the value from the env that is passed in
		if gettingEnv == "" {        // checks to see if it's set or not
			log.Printf("Warning: %s environment variable not set", dns)
		}
		return gettingEnv //if the env is there then return it
	}

	// when it returns the gettingEnv, is it returning it to cloud sql itself?

	var ( //this sends the env to the default func to check if its there -- why is dbUser the only one to not have mustGetenv?
		dbUser         = os.Getenv("DB_USER") //postgres
		dbPwd          = mustGetenv("DB_PWD")
		dbName         = mustGetenv("DB_NAME")              //postgres
		unixSocketPath = mustGetenv("INSTANCE_UNIX_SOCKET") // /cloudsql/jessie-apprentice:us-central1:quotes-database
	)

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s", dbUser, dbPwd, dbName, unixSocketPath) //this creates the path to print and show what database is being used, gives credentials

	//dbPool is the pool of database connections
	var err error

	db, err = sql.Open("pgx", dbURI) //opens a database driver specific to the database being used. in this case, we are using the pgx engine and the path is the dbURI
	if err != nil {                  //if it isn't correct or there, throw error
		return fmt.Errorf("sql.Open: %v", err)
	}
	return err
}

func manageHeader(c *gin.Context) bool { //this manages the authentication api key to ensure it equals cocktailsauce
	headers := c.Request.Header            // this requests the list of headers to read it
	header, exists := headers["X-Api-Key"] //this looks for the specific key we want
	fmt.Println(header)                    //why do we need to print it?

	if exists { //makes sure the right key exists
		if header[0] == "COCKTAILSAUCE" {
			return true
		}
	}
	return false
}

func deleteQuote(c *gin.Context) {
	id := c.Param("id") // what is the purpose of this?
	row := db.QueryRow(fmt.Sprintf("delete from quotes where id = '%s'", id))
	q := &quote{}
	err := row.Scan(&q.ID, &q.Quote, &q.Author)

	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusNoContent, q)
}

func getRandomQuoteSQL(c *gin.Context) {
	if manageHeader(c) {
		row := db.QueryRow("select id, phrase, author from quotes order by RANDOM() limit 1")
		q := &quote{}
		err := row.Scan(&q.ID, &q.Quote, &q.Author)

		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, q)
		return
	} else if !manageHeader(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
	}
}

func getQuoteByIdSQL(c *gin.Context) {
	if manageHeader(c) {
		id := c.Param("id")
		row := db.QueryRow(fmt.Sprintf("select id, phrase, author from quotes where id = '%s'", id))
		q := &quote{}
		err := row.Scan(&q.ID, &q.Quote, &q.Author)

		if err != nil {
			log.Println(err)
		}

		if q.ID == "" {
			c.JSON(http.StatusNotFound, gin.H{"message": "invalid ID"})
			return
		} else {
			c.JSON(http.StatusOK, q)
		}

	} else if !manageHeader(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
	}
}

func addQuoteSQL(c *gin.Context) {
	if manageHeader(c) {
		var newID ID
		q := &quote{}

		newID.ID = uuid.New().String()

		if flaw := c.BindJSON(&q); flaw != nil {
			return
		}

		if len(q.Quote) < 3 || len(q.Author) < 3 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
			return

		} else {
			sqlStatement := `insert into quotes (id, phrase, author) values ($1, $2, $3)`
			_, err := db.Exec(sqlStatement, &newID.ID, &q.Quote, &q.Author)

			if err != nil {
				log.Println(err)
			}

			c.JSON(http.StatusCreated, newID)
		}

	} else if !manageHeader(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
	}
}
