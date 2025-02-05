package interceptor

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

func LoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		startTime := time.Now()
		resp, err = handler(ctx, req)
		duration := time.Since(startTime)

		logx.WithContext(ctx).Infow("RPC调用日志",
			logx.Field("method", info.FullMethod),
			logx.Field("duration", duration),
			logx.Field("error", err),
		)

		return resp, err
	}
}
