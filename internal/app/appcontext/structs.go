package appcontext

import (
	"context"
	"github.com/sacOO7/gowebsocket"
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/harunnryd/btrade/internal/app/config"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

// AppContext ...
type AppContext struct {
	Ctx       context.Context
	CtxCancel context.CancelFunc
	Eg        *errgroup.Group
	AppErrors []error
	Logger    zerolog.Logger
	Config    config.Config
	IsLive    bool

	Olymptrade     gowebsocket.Socket
	OlymptradeOnce sync.Once
	RedisPool      *redis.Pool
}
