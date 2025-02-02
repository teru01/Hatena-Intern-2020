package grpc

import (
	"context"
	"net/http"

	"github.com/hatena/Hatena-Intern-2020/services/fetcher/fetcher"
	pb "github.com/hatena/Hatena-Intern-2020/services/fetcher/pb/fetcher"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server は pb.FetcherServer に対する実装
type Server struct {
	pb.UnimplementedFetcherServer
	healthpb.UnimplementedHealthServer
}

// NewServer は gRPC サーバーを作成する
func NewServer() *Server {
	return &Server{}
}

// Fetcher は受け取った文書を HTML に変換する
func (s *Server) Fetch(ctx context.Context, in *pb.FetchRequest) (*pb.FetchReply, error) {
	title, err := fetcher.Fetch(ctx, in.Uri, &http.Client{})
	if err != nil {
		return nil, err
	}
	return &pb.FetchReply{Title: title}, nil
}
