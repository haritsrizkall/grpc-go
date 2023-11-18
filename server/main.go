package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/haritsrizkall/grpc-go/proto"
	"google.golang.org/grpc"
)

type songService struct {
	pb.SongServiceServer
}

func (s *songService) GetSong(ctx context.Context, req *pb.SongRequest) (*pb.SongResponse, error) {
	return &pb.SongResponse{
		Song: &pb.Song{
			Title:  req.Title,
			Artist: "Test Artist",
			Album:  "Test Album",
		},
	}, nil
}

func main() {
	fmt.Println("Server is running...")
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterSongServiceServer(s, &songService{})
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
