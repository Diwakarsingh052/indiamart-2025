package main

import (
	"fmt"
	"indiamart-2025/6-structs/database"
)

func main() {
	// create new functions to create instances of the struct
	// field of conn struct are not exported
	// which means it can't be modified here
	c := database.NewConn("postgres")
	fmt.Println(c)
	c.Insert()

	//log.New()
	//os.OpenFile()
}
