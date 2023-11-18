package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/haritsrizkall/grpc-go/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client is running...")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:8081", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewSongServiceClient(cc)
	request := &pb.SongRequest{Title: "Radiohead"}

	resp, err := client.GetSong(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive response => %s ", resp.Song.Title)
}
