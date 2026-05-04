package grpcconn

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

const (
	MaxRecvMsgSize = 12 * 1024 * 1024
	timeKeepalive  = 30 * time.Second
)

func New(target string) (*grpc.ClientConn, error) {
	keepaliveParams := keepalive.ClientParameters{
		Time:                timeKeepalive,
		Timeout:             10 * time.Second,
		PermitWithoutStream: false,
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepaliveParams),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxRecvMsgSize)),
	}

	return grpc.NewClient(target, opts...)
}