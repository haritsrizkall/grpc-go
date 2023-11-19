package song

import (
	"context"
	"fmt"

	pb "github.com/haritsrizkall/grpc-go/proto"
)

var songs []*pb.Song

type SongService struct {
	pb.SongServiceServer
}

func (s *SongService) GetSong(ctx context.Context, req *pb.SongRequest) (*pb.SongResponse, error) {
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

func (s *SongService) GetAllSongs(ctx context.Context, req *pb.EmptyRequest) (*pb.SongsResponse, error) {
	stream := &pb.SongsResponse{}
	stream.Songs = songs
	return stream, nil
}

func (s *SongService) AddSong(ctx context.Context, req *pb.Song) (*pb.SongResponse, error) {
	songs = append(songs, req)
	return &pb.SongResponse{
		Song: req,
	}, nil
}

func (s *SongService) DeleteSong(ctx context.Context, req *pb.DeleteSongRequest) (*pb.SongResponse, error) {
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
