package main

import "fmt"

type User struct {
	name  string
	email string
}

type BookAuthor struct {
	// with embedding we get automatic implementation of interface
	// if User implements the interface, then BookAuthor will also implement the interface in case of embedding
	User // embedding // anonymous field // field name would be same as the struct name
	//u     User // not embedding
	books []string
}

func (u *User) UpdateEmail(email string) {
	u.email = email
}

func (b *BookAuthor) AddBook(book string) {
	b.books = append(b.books, book)
}

func (b *BookAuthor) UpdateEmail(email string) {
	b.email = email
}

func main() {

	a := BookAuthor{
		User: User{
			name:  "John",
			email: "",
		},
		books: []string{"Go", "Python"},
	}
	a.User.UpdateEmail("john@email.com")
	a.AddBook("C++")
	fmt.Println(a)

}
