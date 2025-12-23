// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

//

// define a book type with a title and author
type book struct {
	title  string
	author author
}

//

// define a library type to track a list of books
type library struct {
	books []book
}

// define addBook to add a book to the library
func (l *library) addBook(b book) {
	l.books = append(l.books, b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(name string) []book {
	// var results []book
	results := make([]book, 0)
	for _, b := range l.books {
		if b.author.name == name {
			results = append(results, b)
		}

	}
	return results
}

func main() {
	// create a new library
	lib := library{}

	// add 2 books to the library by the same author
	lib.addBook(book{
		title:  "Book One",
		author: author{name: "Author One"},
	})
	lib.addBook(book{
		title:  "Book Two",
		author: author{name: "Author One"},
	})

	// add 1 book to the library by a different author
	lib.addBook(book{
		title:  "Book Three",
		author: author{name: "Author Two"},
	})

	// dump the library with spew
	//fmt.Printf("%#v\n", lib)
	spew.Dump(lib)

	// lookup books by known author in the library
	booksByAuthorOne := lib.lookupByAuthor("Author One")

	// print out the number of books found
	fmt.Printf("Found %d books by Author One\n", len(booksByAuthorOne))

	// print out the first book's title and its author's name
	fmt.Printf("First book: %q by %q\n", booksByAuthorOne[0].title, booksByAuthorOne[0].author.name)

}
