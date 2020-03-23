package model

import "fmt"

//go:generate sh -c "~/code/src/gitlab.com/hookactions/gqlgen-relay/gqlgen-relay -pkg model -name Book -type *Book -cursor > graph/model/book_relay.go"

func (b *Book) GetID() fmt.Stringer {
	return b
}

func (b *Book) String() string {
	return b.ID
}
