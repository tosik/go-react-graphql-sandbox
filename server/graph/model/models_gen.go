// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Book struct {
	ID    string      `json:"id"`
	Title string      `json:"title"`
	Price int         `json:"price"`
	Foo   interface{} `json:"foo"`
}

type NewBook struct {
	Title string      `json:"title"`
	Price int         `json:"price"`
	Foo   interface{} `json:"foo"`
}
