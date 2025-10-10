package main

import (
	"fmt"
	"proto-buf-basics/proto"
)

//https://grpc.io/docs/languages/go/quickstart/
//https://protobuf.dev/programming-guides/proto3/

func main() {
	req := proto.BlogRequest{
		BlogId:  11,
		Foo:     "foo",
		Title:   "title",
		Content: "content",
	}

	fmt.Println(req.GetBlogId())
}
