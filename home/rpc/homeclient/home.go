// Code generated by goctl. DO NOT EDIT!
// Source: home.proto

package homeclient

import (
	"context"

	"microShop/home/rpc/types/home"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BannerReq      = home.BannerReq
	CommonResply   = home.CommonResply
	RankingListReq = home.RankingListReq
	RecommendReq   = home.RecommendReq

	Home interface {
		Banner(ctx context.Context, in *BannerReq, opts ...grpc.CallOption) (*CommonResply, error)
		Recommend(ctx context.Context, in *RecommendReq, opts ...grpc.CallOption) (*CommonResply, error)
		RankingList(ctx context.Context, in *RankingListReq, opts ...grpc.CallOption) (*CommonResply, error)
	}

	defaultHome struct {
		cli zrpc.Client
	}
)

func NewHome(cli zrpc.Client) Home {
	return &defaultHome{
		cli: cli,
	}
}

func (m *defaultHome) Banner(ctx context.Context, in *BannerReq, opts ...grpc.CallOption) (*CommonResply, error) {
	client := home.NewHomeClient(m.cli.Conn())
	return client.Banner(ctx, in, opts...)
}

func (m *defaultHome) Recommend(ctx context.Context, in *RecommendReq, opts ...grpc.CallOption) (*CommonResply, error) {
	client := home.NewHomeClient(m.cli.Conn())
	return client.Recommend(ctx, in, opts...)
}

func (m *defaultHome) RankingList(ctx context.Context, in *RankingListReq, opts ...grpc.CallOption) (*CommonResply, error) {
	client := home.NewHomeClient(m.cli.Conn())
	return client.RankingList(ctx, in, opts...)
}
