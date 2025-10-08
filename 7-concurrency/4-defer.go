package main

import (
	"fmt"
)

func main() {
	// defer statements are executed in when the function returns
	// defer maintains a stack
	// first in last out
	// defer gives guarantee that the function will be executed even in case of panic
	defer fmt.Println(1)
	defer fmt.Println(2)
	//panic("panic")
	// don't use os.Exit() other than startup code
	// os.Exit() will not let  defer statements to execute
	//os.Exit(1)
	//log.Fatal() // internally calls os.Exit(1)
	fmt.Println(3)
	fmt.Println(4)

}
