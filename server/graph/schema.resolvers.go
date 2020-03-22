package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"io"
	"log"
  "math/rand"
  "fmt"

	"github.com/tosik/go-react-graphql-sandbox/server/graph/generated"
	"github.com/tosik/go-react-graphql-sandbox/server/graph/model"
)

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
  newBook := model.Book {
    ID : fmt.Sprint(rand.Int()),
    Title : input.Title,
    Price : input.Price,
    Foo : input.Foo,
  }
  // err := r.Resolver.Coll.Actions().Put(&input).Get(&got).Do(ctx)
  err := r.Resolver.Coll.Put(ctx, &newBook)
  if err != nil {
    log.Fatalln(err)
    return nil, err
  }

  got := &model.Book{ ID: newBook.ID }
  err = r.Resolver.Coll.Get(ctx, got)
  if err != nil {
    log.Fatalln(err)
    return nil, err
  }

  return got, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	iter := r.Resolver.Coll.Query().Get(ctx)
	defer iter.Stop()

  dest := []*model.Book{}
	for {
		var book model.Book
		err := iter.Next(ctx, &book)
		if err == io.EOF {
			break
		} else if err != nil {
      log.Fatalln(err)
			return nil, err
		}

		dest = append(dest, &book)
	}

	return dest, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
