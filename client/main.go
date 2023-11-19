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

	songClient := pb.NewSongServiceClient(cc)
	lyricClient := pb.NewLyricServiceClient(cc)

	// AddSong
	song := &pb.Song{
		Title:  "Lemon Tree",
		Artist: "Fools Garden",
		Album:  "Dish of the Day",
	}
	resp, err := songClient.AddSong(context.Background(), song)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Song)

	// add lyric
	lyric1 := &pb.Lyric{
		Id:        "1",
		SongTitle: resp.Song.Title,
		Content:   "I'm sitting here in a boring room",
		Duration:  3,
		Order:     1,
	}
	_, err = lyricClient.AddLyric(context.Background(), lyric1)
	if err != nil {
		log.Fatal(err)
	}

	lyric2 := &pb.Lyric{
		Id:        "2",
		SongTitle: resp.Song.Title,
		Content:   "I'm sitting here in a boring room 2",
		Duration:  10,
		Order:     2,
	}
	_, err = lyricClient.AddLyric(context.Background(), lyric2)
	if err != nil {
		log.Fatal(err)
	}

	lyric3 := &pb.Lyric{
		Id:        "3",
		SongTitle: resp.Song.Title,
		Content:   "I'm sitting here in a boring room 3",
		Duration:  1,
		Order:     3,
	}
	_, err = lyricClient.AddLyric(context.Background(), lyric3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Lyrics added")
	fmt.Println("Lyrics 1", lyric1)
	fmt.Println("Lyrics 2", lyric2)

	// stream lyrics
	stream, err := lyricClient.GetSongLyrics(context.Background(), &pb.GetSongLyricsRequest{
		SongTitle: resp.Song.Title,
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		lyric, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(lyric.Content)
	}
}
