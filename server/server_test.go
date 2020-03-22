package main

import (
	"context"
	"testing"

	"gocloud.dev/docstore"
)

func TestMain(t *testing.T) {
	url := "mem://test/ID"
	ctx := context.Background()
	coll, err := docstore.OpenCollection(ctx, url)
	defer coll.Close()
	if err != nil {
		t.Fatalf("cannot open memstore: %s", err)
	}

	err = putSampleData(ctx, coll)
	if err != nil {
		t.Fatalf("cannot put sample: %s", err)
	}
}
