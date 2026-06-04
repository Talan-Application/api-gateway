package grpcconn

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"

	"github.com/Talan-Application/api-gateway/internal/ctxkeys"
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
		grpc.WithUnaryInterceptor(authForwardInterceptor),
	}

	return grpc.NewClient(target, opts...)
}

func authForwardInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	if header, ok := ctx.Value(ctxkeys.AuthHeaderKey).(string); ok && header != "" {
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", header)
	}
	return invoker(ctx, method, req, reply, cc, opts...)
}
