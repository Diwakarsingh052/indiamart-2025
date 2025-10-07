package main

// moduleName/pkgName
import (
	"fmt"
	"learn-pkg/calc"
	"learn-pkg/db"
)

func main() {

	calc.Add(1, 2)
	db.NewConnection("postgres")
	fmt.Println("Inserting in db", db.GetConnection())
}
