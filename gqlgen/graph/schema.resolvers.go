package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorrelljd21/quotes-starter/gqlgen/graph/generated"
	"github.com/gorrelljd21/quotes-starter/gqlgen/graph/model"
)

// Quote is the resolver for the quote field.
func (r *queryResolver) Quote(ctx context.Context) (*model.Quote, error) {
	oneQuote := &model.Quote{
		ID:     "321",
		Quote:  "hello",
		Author: "me",
	}

	return oneQuote, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
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
func (r *queryResolver) Quotes(ctx context.Context) ([]*model.Quote, error) {
	panic(fmt.Errorf("not implemented: Quotes - quotes"))
}
func (r *queryResolver) ID(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}
