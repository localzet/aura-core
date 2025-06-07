package singbridge

import (
	"context"

	"github.com/localzet/aura/common/errors"
	"github.com/sagernet/sing/common/logger"
)

var _ logger.ContextLogger = (*AuraLogger)(nil)

type AuraLogger struct {
	newError func(values ...any) *errors.Error
}

func NewLogger(newErrorFunc func(values ...any) *errors.Error) *AuraLogger {
	return &AuraLogger{
		newErrorFunc,
	}
}

func (l *AuraLogger) Trace(args ...any) {
}

func (l *AuraLogger) Debug(args ...any) {
	errors.LogDebug(context.Background(), args...)
}

func (l *AuraLogger) Info(args ...any) {
	errors.LogInfo(context.Background(), args...)
}

func (l *AuraLogger) Warn(args ...any) {
	errors.LogWarning(context.Background(), args...)
}

func (l *AuraLogger) Error(args ...any) {
	errors.LogError(context.Background(), args...)
}

func (l *AuraLogger) Fatal(args ...any) {
}

func (l *AuraLogger) Panic(args ...any) {
}

func (l *AuraLogger) TraceContext(ctx context.Context, args ...any) {
}

func (l *AuraLogger) DebugContext(ctx context.Context, args ...any) {
	errors.LogDebug(ctx, args...)
}

func (l *AuraLogger) InfoContext(ctx context.Context, args ...any) {
	errors.LogInfo(ctx, args...)
}

func (l *AuraLogger) WarnContext(ctx context.Context, args ...any) {
	errors.LogWarning(ctx, args...)
}

func (l *AuraLogger) ErrorContext(ctx context.Context, args ...any) {
	errors.LogError(ctx, args...)
}

func (l *AuraLogger) FatalContext(ctx context.Context, args ...any) {
}

func (l *AuraLogger) PanicContext(ctx context.Context, args ...any) {
}
