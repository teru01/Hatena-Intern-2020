package converter

import (
	"context"

	pb_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	"google.golang.org/grpc"
)

type DummyFetchClient struct {
	callCount int
}

func (d *DummyFetchClient) Fetch(ctx context.Context, in *pb_fetcher.FetchRequest, opts ...grpc.CallOption) (*pb_fetcher.FetchReply, error) {
	d.callCount++
	return &pb_fetcher.FetchReply{Title: "success"}, nil
}

func (d *DummyFetchClient) CallCount() int {
	return d.callCount
}
