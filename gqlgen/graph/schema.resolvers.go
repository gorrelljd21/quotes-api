package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

	if err != nil {
		return nil, err
	}

	bufferResponse := bytes.NewBuffer(response)

	stringKey := ctx.Value("API-Key").(string)

	request, err := http.NewRequest("POST", "http://34.160.90.176:80/quote", bufferResponse)
	request.Header.Set("X-Api-Key", stringKey)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, _ := client.Do(request)

	switch resp.StatusCode {
	case 404:
		return nil, errors.New("invalid input")
	case 401:
		return nil, errors.New("unauthorized")
	}

	if len(input.Author) < 3 || len(input.Quote) < 3 {
		return nil, errors.New("invalid input")
	}

	otherResponse, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(otherResponse, quote)

	return quote, nil
}

// DeleteQuote is the resolver for the deleteQuote field.
func (r *mutationResolver) DeleteQuote(ctx context.Context, id string) (*model.DeleteQuote, error) {

	//check if the quote exists
	_, err := r.Query().QuoteID(ctx, id)

	if err != nil {
		return nil, err
	}

	//form request to REST api
	stringKey := ctx.Value("API-Key").(string)
	request, err := http.NewRequest("DELETE", fmt.Sprintf("http://34.160.90.176:80/quote/%s", id), nil)
	request.Header.Set("X-Api-Key", stringKey)

	if err != nil {
		return nil, err
	}

	//deleting the id (firing off the request)
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	// shows success message
	deleteQuote := &model.DeleteQuote{
		Code:    resp.StatusCode,
		Message: "successfully deleted",
	}

	//checks for unauthorized
	switch resp.StatusCode {
	case 401:
		return nil, errors.New("unauthorized")
	}

	return deleteQuote, nil
}

// Quote is the resolver for the quote field.
func (r *queryResolver) Quote(ctx context.Context) (*model.Quote, error) {
	var randQuote *model.Quote

	stringKey := ctx.Value("API-Key").(string)

	request, err := http.NewRequest("GET", "http://34.160.90.176:80/quote", nil)
	request.Header.Set("X-Api-Key", stringKey)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, _ := client.Do(request)

	switch resp.StatusCode {
	case 401:
		return nil, errors.New("unauthorized")
	}

	requestBody, err := io.ReadAll(resp.Body)

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
	stringKey := ctx.Value("API-Key").(string)

	request, err := http.NewRequest("GET", fmt.Sprintf("http://34.160.90.176:80/quote/%s", id), nil)
	request.Header.Set("X-Api-Key", stringKey)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, _ := client.Do(request)

	switch resp.StatusCode {
	case 404:
		return nil, errors.New("id not found")
	case 401:
		return nil, errors.New("unauthorized")
	}

	var quoteById *model.Quote
	requestBody, err := io.ReadAll(resp.Body)
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
