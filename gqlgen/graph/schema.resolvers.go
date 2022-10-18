package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorrelljd21/quotes-starter/gqlgen/graph/generated"
	"github.com/gorrelljd21/quotes-starter/gqlgen/graph/model"
)

// AddQuote is the resolver for the addQuote field
func (r *mutationResolver) InsertQuote(ctx context.Context, input model.NewQuote) (*model.Quote, error) {
	quote := &model.Quote{
		Quote:  input.Quote,
		Author: input.Author,
	}

	response, err := json.Marshal(&quote)
	bufferResponse := bytes.NewBuffer(response)

	request, err := http.NewRequest("POST", "http://34.160.90.176:80/quote", bufferResponse)
	request.Header.Set("X-Api-Key", "COCKTAILSAUCE")
	// request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, _ := client.Do(request)

	otherResponse, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(otherResponse, quote)

	return quote, nil
}

// DeleteQuote is the resolver for the deleteQuote field.
func (r *mutationResolver) DeleteQuote(ctx context.Context, id string) (*model.Quote, error) {
	panic(fmt.Errorf("not implemented: DeleteQuote - deleteQuote"))
}

// Quote is the resolver for the quote field.
func (r *queryResolver) Quote(ctx context.Context) (*model.Quote, error) {
	var randQuote *model.Quote

	request, err := http.NewRequest("GET", "http://34.160.90.176:80/quote", nil)
	request.Header.Set("X-Api-Key", "COCKTAILSAUCE")

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
	request, err := http.NewRequest("GET", fmt.Sprintf("http://34.160.90.176:80/quote/%s", id), nil)
	request.Header.Set("X-Api-Key", "COCKTAILSAUCE")

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
