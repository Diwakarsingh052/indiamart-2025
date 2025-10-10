package main

import "testing"

// // the below command generates memory profile
// // go test -run none -bench . -benchtime 3s -benchmem -v
func BenchmarkAnalyzeText(b *testing.B) {
	for b.Loop() {
		_, err := AnalyzeText("book.txt")
		if err != nil {
			b.Error(err)
		}
	}
}
