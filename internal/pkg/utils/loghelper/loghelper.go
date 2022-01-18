package loghelper

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

type LogZero struct {
	Logz   zerolog.Logger
	LogCtx context.Context
}

func (l *LogZero) Printf(ctx context.Context, format string, v ...interface{}) {
	startTime := ctx.Value("start_time")
	uptime := ""
	if startTime != nil {
		uptime = time.Since(startTime.(time.Time)).String()
	}

	ctx = l.Logz.WithContext(l.LogCtx)
	zerolog.Ctx(ctx).
		Info().
		Str("uptime", uptime).
		Msg(fmt.Sprintf(format, v...))
}

type LogCaller struct {
	File     string
	Line     int
	Function string
}

func New() LogCaller {
	pc, f, l, _ := runtime.Caller(4)
	return LogCaller{
		File:     f,
		Line:     l,
		Function: runtime.FuncForPC(pc).Name(),
	}
}

func (lc LogCaller) MarshalZerologObject(e *zerolog.Event) {
	e.Str("file", lc.File).
		Int("line", lc.Line).
		Str("function", lc.Function)
}

func AddStrWithKey(ctx context.Context, key, msg string) {
	zerolog.Ctx(ctx).UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str(key, msg)
	})
}

func AddStr(ctx context.Context, msg string) {
	zerolog.Ctx(ctx).UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str(strconv.FormatInt(time.Now().UnixNano(), 10), msg)
	})
}

func AddErr(ctx context.Context, err error) {
	if err != nil {
		zerolog.Ctx(ctx).UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Object(strconv.FormatInt(time.Now().UnixNano(), 10), New()).
				AnErr(strconv.FormatInt(time.Now().UnixNano(), 10), err)
		})
	}
}

func AddErrAndStr(ctx context.Context, msg string, err error) {
	zerolog.Ctx(ctx).UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Object(strconv.FormatInt(time.Now().UnixNano(), 10), New()).
			AnErr(strconv.FormatInt(time.Now().UnixNano(), 10), err).
			Str(strconv.FormatInt(time.Now().UnixNano(), 10), msg)
	})
}

func AddStrOrPanic(ctx context.Context, err error, strFatal string, strInfo string) {
	startTime := ctx.Value("start_time")
	uptime := ""
	if startTime != nil {
		uptime = time.Since(startTime.(time.Time)).String()
	}

	if err != nil {
		zerolog.Ctx(ctx).
			Panic().
			Err(err).
			Caller(2).
			Str("uptime", uptime).
			Msg(strFatal)
	} else {
		zerolog.Ctx(ctx).UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str(strconv.FormatInt(time.Now().UnixNano(), 10), strInfo)
		})
	}
}

func AddStrOrAddErr(ctx context.Context, err error, strError string, strInfo string) {
	if err != nil {
		zerolog.Ctx(ctx).UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Object(strconv.FormatInt(time.Now().UnixNano(), 10), New()).
				AnErr(strconv.FormatInt(time.Now().UnixNano(), 10), err).
				Str(strconv.FormatInt(time.Now().UnixNano(), 10), strError)
		})
	} else {
		zerolog.Ctx(ctx).UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str(strconv.FormatInt(time.Now().UnixNano(), 10), strInfo)
		})
	}
}

func Panic(ctx context.Context, msg string, err error) {
	startTime := ctx.Value("start_time")
	uptime := ""
	if startTime != nil {
		uptime = time.Since(startTime.(time.Time)).String()
	}

	zerolog.Ctx(ctx).
		Panic().
		Err(err).
		Caller(2).
		Str("uptime", uptime).
		Msg(msg)
}

func MsgOrPanic(ctx context.Context, msg string, err error) {
	startTime := ctx.Value("start_time")
	uptime := ""
	if startTime != nil {
		uptime = time.Since(startTime.(time.Time)).String()
	}

	if err != nil {
		zerolog.Ctx(ctx).
			Panic().
			Err(err).
			Caller(2).
			Str("uptime", uptime).
			Msg(msg)
	} else {
		zerolog.Ctx(ctx).
			Info().
			Str("uptime", uptime).
			Msg(msg)
	}
}

func Msg(ctx context.Context, msg string) {
	startTime := ctx.Value("start_time")
	uptime := ""
	if startTime != nil {
		uptime = time.Since(startTime.(time.Time)).String()
	}

	zerolog.Ctx(ctx).
		Info().
		Str("uptime", uptime).
		Msg(msg)
}

func Err(ctx context.Context, msg string, err error) {
	startTime := ctx.Value("start_time")
	uptime := ""
	if startTime != nil {
		uptime = time.Since(startTime.(time.Time)).String()
	}
	zerolog.Ctx(ctx).
		Err(err).
		Caller(2).
		Str("uptime", uptime).
		Msg(msg)
}
