package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Names slice with all the provided names
var names = []string{
	"Abhinay",
	"Sindhuja",
	"Chandrashekhar",
	"Rameshwar",
	"Mohit",
	"Nikita",
	"Vishwas Kalra",
	"Manish Solanki",
	"Vidyadhar",
	"Akshat",
	"Samyak",
	"Saurabh",
	"Anwer",
	"Ravideep",
	"Sriram",
	"Anushri",
	"Bhargavi",
	"Satyam",
	"Rahul Sehgal",
	"Satyam Yadav",
	"Govind",
	"Jatin",
	"Roma",
}

// RandomNamePicker picks a random name from a slice of names
func RandomNamePicker(names []string) string {
	if len(names) == 0 {
		return ""
	}

	// Create a new random source with current time as seed
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Pick a random index
	randomIndex := r.Intn(len(names))
	return names[randomIndex]
}

func main() {

	fmt.Println(RandomNamePicker(names))
}
