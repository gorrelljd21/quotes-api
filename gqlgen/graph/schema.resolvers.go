package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// "github.com/gin-gonic/gin"
	"github.com/gorrelljd21/quotes-starter/gqlgen/graph/generated"
	"github.com/gorrelljd21/quotes-starter/gqlgen/graph/model"
)

// Quote is the resolver for the quote field.
func (r *queryResolver) Quote(ctx context.Context) (*model.Quote, error) {
	var randQuote *model.Quote

	response, err := http.Get("http://0.0.0.0:8080/quote")

	if err != nil {
		fmt.Print(err.Error())
	}

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Print(err.Error())
	}

	err = json.Unmarshal(responseBody, &randQuote)

	if err != nil {
		fmt.Print(err.Error())
	}

	return randQuote, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
