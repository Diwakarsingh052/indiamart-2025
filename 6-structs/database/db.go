package database

import "fmt"

type Conn struct {
	db string
}

func NewConn(db string) Conn {
	c := Conn{db: db}
	return c
}

func (c Conn) Insert() {
	fmt.Println("Inserting in db", c.db)
}
