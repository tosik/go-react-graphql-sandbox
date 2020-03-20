package main

import (
	"log"
	"net/http"
	"os"
  "context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tosik/go-react-graphql-sandbox/server/graph"
	"github.com/tosik/go-react-graphql-sandbox/server/graph/generated"

  "gocloud.dev/docstore"
  _ "gocloud.dev/docstore/mongodocstore"
)

type Book struct {
  ID int
  Title string
  Price int
  DocstoreRevision interface{}
}

const defaultPort = "8080"

func main() {
  {
    ctx := context.Background()

    url := "mongo://foo/books?id_field=ID"

    coll, err := docstore.OpenCollection(ctx, url)
    if err != nil {
      log.Fatalln(err)
    }

    books := []Book{
      { ID: 1, Title: "Alice In Wonderland", Price: 123 },
      { ID: 2, Title: "Cinderella", Price: 345 },
    }

    for _, book := range books {
      log.Println(book)
      err = coll.Put(ctx, &book)
      if err != nil {
        log.Fatalln(err)
      }
    }

    defer coll.Close()
  }

  {
    port := os.Getenv("PORT")
    if port == "" {
      port = defaultPort
    }

    srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)

    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
  }
}
