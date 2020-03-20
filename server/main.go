package main

import (
  "log"
  "context"

  "gocloud.dev/docstore"
  _ "gocloud.dev/docstore/mongodocstore"
)

type Book struct {
  ID int
  Title string
  Price int
  DocstoreRevision interface{}
}

func main() {
  ctx := context.Background()

  {
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
}
