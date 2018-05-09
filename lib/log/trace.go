package log

import (
	"context"
	"fmt"
)

const (
	// TraceKey is the key name in context.Context.
	TraceKey CtxKey = "trace"
)

// CtxKey defines a new type in order to avoid value conflicting in context.Context.
type CtxKey string

// Debugfx write one line of debug level log with the providing context, format and arguments.
func Debugfx(ctx context.Context, format string, args ...interface{}) {
	log.Debugf(formatCtx(ctx, format), args...)
}

// Infofx write one line of info level log with the providing context, format and arguments.
func Infofx(ctx context.Context, format string, args ...interface{}) {
	log.Infof(formatCtx(ctx, format), args...)
}

// Noticefx write one line of notice level log with the providing context, format and arguments.
func Noticefx(ctx context.Context, format string, args ...interface{}) {
	log.Noticef(formatCtx(ctx, format), args...)
}

// Warningfx write one line of warning level log with the providing context, format and arguments.
func Warningfx(ctx context.Context, format string, args ...interface{}) {
	log.Warningf(formatCtx(ctx, format), args...)
}

// Errorfx write one line of error level log with the providing context, format and arguments.
func Errorfx(ctx context.Context, format string, args ...interface{}) {
	log.Errorf(formatCtx(ctx, format), args...)
}

// Criticalfx write one line of critical level log with the providing context, format and arguments.
func Criticalfx(ctx context.Context, format string, args ...interface{}) {
	log.Criticalf(formatCtx(ctx, format), args...)
}

// Fatalfx write one line of critical level log with the providing context, format and arguments, and then exit the program.
func Fatalfx(ctx context.Context, format string, args ...interface{}) {
	log.Fatalf(formatCtx(ctx, format), args...)
}

func formatCtx(ctx context.Context, format string) string {
	if ctx != nil {
		return fmt.Sprintf("trace: %s, %s", ctx.Value(TraceKey), format)
	}
	return format
}
