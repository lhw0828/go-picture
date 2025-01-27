package trace

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const (
	TraceIDKey contextKey = "trace_id"
)

func NewContext(ctx context.Context) context.Context {
	traceID := uuid.New().String()
	return context.WithValue(ctx, TraceIDKey, traceID)
}

func GetTraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	traceID, _ := ctx.Value(TraceIDKey).(string)
	return traceID
}