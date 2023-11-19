package main

import (
	"fmt"
	"net"

	pb "github.com/haritsrizkall/grpc-go/proto"
	"github.com/haritsrizkall/grpc-go/server/lyric"
	"github.com/haritsrizkall/grpc-go/server/song"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Server is running...")
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	songSvc := song.SongService{}
	lyricSvc := lyric.LyricService{}
	pb.RegisterSongServiceServer(s, &songSvc)
	pb.RegisterLyricServiceServer(s, &lyricSvc)
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
