package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"

	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/memdocstore"
	_ "gocloud.dev/docstore/mongodocstore"

	"github.com/tosik/go-react-graphql-sandbox/server/graph"
	"github.com/tosik/go-react-graphql-sandbox/server/graph/generated"
	"github.com/tosik/go-react-graphql-sandbox/server/graph/model"
)

//go:generate sh -c "go run gitlab.com/hookactions/gqlgen-relay -pkg model -name Book -type *Book -cursor > graph/model/book_relay.go"

const defaultPort = "8080"

func main() {
	ctx := context.Background()

	// url := "mongo://foo/books?id_field=ID"
	url := "mem://foo/ID"

	coll, err := docstore.OpenCollection(ctx, url)
	defer coll.Close()
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	if err := putSampleData(ctx, coll); err != nil {
		log.Fatalln(err)
	}
	listen(coll)
}

func putSampleData(ctx context.Context, coll *docstore.Collection) error {
	books := []model.Book{
		{ID: "1", Title: "Alice In Wonderland", Price: 123, Foo: 98765},
		{ID: "2", Title: "Cinderella", Price: 345, Foo: "STRING TYPE"},
	}
	for _, book := range books {
		if err := coll.Put(ctx, &book); err != nil {
			return err
		}
	}

	return nil
}

func listen(coll *docstore.Collection) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			Coll: coll,
		}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
