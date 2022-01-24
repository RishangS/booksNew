// models.book.go

package main

type book struct {
	ID     int    `json:"id" bson:"id,omitempty"`
	Title  string `json:"title" bson:"title,omitempty"`
	Author string `json:"author" bson:"author,omitempty"`
}
