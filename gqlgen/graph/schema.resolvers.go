package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorrelljd21/quotes-starter/gqlgen/graph/generated"
	"github.com/gorrelljd21/quotes-starter/gqlgen/graph/model"
)

// AddQuote is the resolver for the addQuote field.
func (r *mutationResolver) AddQuote(ctx context.Context, input model.newQuote) (*model.Quote, error) {
	request, err := http.NewRequest("POST", "http://0.0.0.0:8080/quote", nil)
	request.Header.Set("x-api-key", "COCKTAILSAUCE")

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, _ := client.Do(request)

	quote := &model.Quote{
		ID:     fmt.Sprintf("%v", uuid.NewV4()),
		Quote:  input.Quote,
		Author: input.Author,
	}

	return &quote, nil
}

// Quote is the resolver for the quote field.
func (r *queryResolver) Quote(ctx context.Context) (*model.Quote, error) {
	var randQuote *model.Quote

	request, err := http.NewRequest("GET", "http://0.0.0.0:8080/quote", nil)
	request.Header.Set("x-api-key", "COCKTAILSAUCE")

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, _ := client.Do(request)

	requestBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(requestBody, &randQuote)

	if err != nil {
		return nil, err
	}
	return randQuote, nil
}

// QuoteID is the resolver for the quoteId field.
func (r *queryResolver) QuoteID(ctx context.Context, id string) (*model.Quote, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("http://0.0.0.0:8080/quote/%s", id), nil)
	request.Header.Set("x-api-key", "COCKTAILSAUCE")

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, _ := client.Do(request)

	var quoteById *model.Quote
	requestBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(requestBody, &quoteById)
	if err != nil {
		return nil, err
	}

	return quoteById, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
