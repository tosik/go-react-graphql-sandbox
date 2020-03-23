package graph

import (
	"encoding/base64"
	"math/rand"

	"github.com/vmihailenco/msgpack"
)

func GenerateId() string {
	b, err := msgpack.Marshal(rand.Int())
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
