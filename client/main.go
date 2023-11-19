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
	request := pb.Song{
		Title:  "Lemon Tree",
		Artist: "Fools Garden",
		Album:  "Dish of the Day",
	}
	_, err = client.AddSong(context.Background(), &request)
	if err != nil {
		log.Fatal(err)
	}

	request2 := pb.Song{
		Title:  "Hikayat Cinta",
		Artist: "Glenn Fredly",
		Album:  "Selamat Pagi, Dunia!",
	}
	_, err = client.AddSong(context.Background(), &request2)
	if err != nil {
		log.Fatal(err)
	}

	request3 := pb.Song{
		Title:  "Alkohol",
		Artist: "Sisitipsi",
		Album:  "Sisitipsi ALbum",
	}
	_, err = client.AddSong(context.Background(), &request3)
	if err != nil {
		log.Fatal(err)
	}

	// get all
	songs, err := client.GetAllSongs(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Song 1: %v", songs.Songs[0])
	log.Printf("Song 2: %v", songs.Songs[1])
	log.Printf("Song 3: %v", songs.Songs[2])

	// delete
	song, err := client.DeleteSong(context.Background(), &pb.DeleteSongRequest{Title: "Lemon Tree"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted song: %v", song.Song)

	// get all
	songs, err = client.GetAllSongs(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Song 1: %v", songs.Songs)
}
