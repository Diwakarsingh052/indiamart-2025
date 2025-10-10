package main

import (
	"fmt"
	"log"
	pb "server/gen/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (us *UserService) GetPosts(req *pb.GetPostsRequest, stream grpc.ServerStreamingServer[pb.GetPostsResponse]) error {

	id := req.GetUserId()

	// assume you hit the database and get the posts, maybe 1000s of them
	// and we are going to send the posts in chunks
	log.Println("GetPosts: ", "fetching all posts for user id ", id)

	// assuming we got some batches from the db

	batch1 := []*pb.Post{
		{
			Title:  "The Science of Design",
			Author: "Author 1",
			Body:   "Body of post 1",
		},
		{
			Title:  "The Politics of Power",
			Author: "Author 2",
			Body:   "Body of post 2",
		},
		{
			Title:  "The Art of Programming",
			Author: "Author 3",
			Body:   "Body of post 3",
		},
	}

	err := stream.Send(&pb.GetPostsResponse{Posts: batch1})
	if err != nil {
		return status.Errorf(codes.Internal, "error sending posts")
	}

	log.Println("sent first batch for user id", id)

	time.Sleep(5 * time.Second)

	//constructing the second batch
	batch2 := []*pb.Post{
		{
			Title:  "Post 11",
			Author: "Author 1",
			Body:   "Body of post 1",
		},
		{
			Title:  "Post 21",
			Author: "Author 2",
			Body:   "Body of post 2",
		},
		{
			Title:  "Post 31",
			Author: "Author 3",
			Body:   "Body of post 3",
		},
	}

	res := pb.GetPostsResponse{Posts: batch2}
	err = stream.Send(&res)
	if err != nil {
		return status.Errorf(codes.Internal, "error sending posts")

	}
	log.Println("sent second batch for user id", id)
	fmt.Println()
	log.Println("all posts are sent for user id", id)
	fmt.Println()

	return nil

}
