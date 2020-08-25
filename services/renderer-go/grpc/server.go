package grpc

import (
	"context"

	"github.com/hatena/Hatena-Intern-2020/services/renderer-go/converter"
	pb "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/renderer"
	"github.com/hatena/Hatena-Intern-2020/services/renderer-go/renderer"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server は pb.RendererServer に対する実装
type Server struct {
	pb.UnimplementedRendererServer
	healthpb.UnimplementedHealthServer
	lineConverters []converter.LineConverter
	wholeConverter []converter.WholeConverter
}

// NewServer は gRPC サーバーを作成する
func NewServer(lc []converter.LineConverter, wc []converter.WholeConverter) *Server {
	return &Server{
		lineConverters: lc,
		wholeConverter: wc,
	}
}

// Render は受け取った文書を HTML に変換する
func (s *Server) Render(ctx context.Context, in *pb.RenderRequest) (*pb.RenderReply, error) {
	html, err := renderer.Render(ctx, in.Src, s.lineConverters, s.wholeConverter)
	if err != nil {
		return nil, err
	}
	return &pb.RenderReply{Html: html}, nil
}
