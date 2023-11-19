package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/haritsrizkall/grpc-go/proto"
	"google.golang.org/grpc"
)

var songs []*pb.Song

type songService struct {
	pb.SongServiceServer
}

func (s *songService) GetSong(ctx context.Context, req *pb.SongRequest) (*pb.SongResponse, error) {
	var song *pb.Song
	for _, s := range songs {
		if s.Title == req.Title {
			song = s
			break
		}
	}
	if song == nil {
		return nil, fmt.Errorf("song not found: %v", req.Title)
	}
	return &pb.SongResponse{
		Song: song,
	}, nil
}

func (s *songService) GetAllSongs(ctx context.Context, req *pb.EmptyRequest) (*pb.SongsResponse, error) {
	stream := &pb.SongsResponse{}
	stream.Songs = songs
	return stream, nil
}

func (s *songService) AddSong(ctx context.Context, req *pb.Song) (*pb.SongResponse, error) {
	songs = append(songs, req)
	return &pb.SongResponse{
		Song: req,
	}, nil
}

func (s *songService) DeleteSong(ctx context.Context, req *pb.DeleteSongRequest) (*pb.SongResponse, error) {
	var idx int
	for i, s := range songs {
		if s.Title == req.Title {
			idx = i
			break
		}
	}
	songs = append(songs[:idx], songs[idx+1:]...)
	return &pb.SongResponse{
		Song: &pb.Song{},
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
