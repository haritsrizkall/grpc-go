package lyric

import (
	"context"
	"fmt"
	"time"

	pb "github.com/haritsrizkall/grpc-go/proto"
)

var lyrics []*pb.Lyric

type LyricService struct {
	pb.LyricServiceServer
}

func (s *LyricService) AddLyric(ctx context.Context, req *pb.Lyric) (*pb.Lyric, error) {
	lyrics = append(lyrics, req)
	return req, nil
}

func (s *LyricService) GetSongLyrics(req *pb.GetSongLyricsRequest, stream pb.LyricService_GetSongLyricsServer) error {
	lyricsData := []*pb.Lyric{}
	for _, l := range lyrics {
		if l.SongTitle == req.SongTitle {
			lyricsData = append(lyricsData, l)
		}
	}
	fmt.Println("Sending lyrics... ", lyricsData)
	for _, l := range lyricsData {
		if err := stream.Send(l); err != nil {
			return err
		}
		time.Sleep(time.Duration(l.Duration) * time.Second)
	}
	return nil
}

func (s *LyricService) UpdateLyric(ctx context.Context, req *pb.UpdateLyricRequest) (*pb.Lyric, error) {
	var lyric *pb.Lyric
	for _, l := range lyrics {
		if l.SongTitle == req.SongTitle {
			lyric = l
			break
		}
	}
	if lyric == nil {
		return nil, nil
	}
	lyric.Content = req.Content
	return lyric, nil
}
