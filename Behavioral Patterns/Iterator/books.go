package main

import "fmt"

type BookType int

const (
	HardCover BookType = iota
	SoftCover
	PaperBack
	Ebook
)

type Book struct {
	Name      string
	Author    string
	PageCount int
	Type      BookType
}
type Library struct {
	Collection []Book
}

func (l *Library) IterateBooks(f func(Book) error) {
	var err error
	for _, b := range l.Collection {
		err = f(b)
		if err != nil {
			fmt.Println("Error encountered:", err)
			break
		}
	}

}

func (l *Library) createIterator() iterator {
	return &BookIterator{
		books: l.Collection,
	}
}

var lib *Library = &Library{
	Collection: []Book{
		{
			Name:      "War and Peace",
			Author:    "Leo Tolstoy",
			PageCount: 864,
			Type:      HardCover,
		}, {
			Name:      "Crime and Punishment",
			Author:    "Fyodor Dostoevsky",
			PageCount: 900,
			Type:      SoftCover,
		}, {
			Name:      "Brave New World",
			Author:    "Aldous Huxley",
			PageCount: 300,
			Type:      SoftCover,
		},
		{
			Name:      "Thus Spoke Zarathustra",
			Author:    "Friedrich Nietzsche",
			PageCount: 1023,
			Type:      HardCover,
		}, {
			Name:      "The Trial",
			Author:    "Franz Kafka",
			PageCount: 483,
			Type:      Ebook,
		}, {
			Name:      "Nausea",
			Author:    "Jean Paul Sartre",
			PageCount: 543,
			Type:      PaperBack,
		},
	},
}
