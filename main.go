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
	r.GET("/quotes", getRandomQuoteSQL)
	r.GET("/quotes/:id", getQuoteByIdSQL)
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

		sqlStatement := `insert into quotes (id, phrase, author) values ($1, $2, $3)`
		_, err := db.Exec(sqlStatement, &newID.ID, &q.Quote, &q.Author)
		if err != nil {
			log.Println(err)
		}

		if len(q.Quote) < 3 || len(q.Author) < 3 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
			return
		}
		c.JSON(http.StatusCreated, newID)
	} else if !manageHeader(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
	}
}
