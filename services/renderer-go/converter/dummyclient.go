package converter

import (
	"context"

	pb_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	"google.golang.org/grpc"
)

type DummyFetchClient struct{}

func (d *DummyFetchClient) Fetch(ctx context.Context, in *pb_fetcher.FetchRequest, opts ...grpc.CallOption) (*pb_fetcher.FetchReply, error) {
	return &pb_fetcher.FetchReply{Title: "success"}, nil
}
