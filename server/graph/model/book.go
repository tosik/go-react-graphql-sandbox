package model

import "fmt"

func (b *Book) GetID() fmt.Stringer {
	return b
}

func (b *Book) String() string {
	return b.ID
}
