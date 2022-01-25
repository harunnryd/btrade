package config

import (
	"github.com/harunnryd/btrade/internal/app/config/core"
	"github.com/harunnryd/btrade/internal/app/config/core/brokers/v1.0/olymptrade"
)

// Config ...
type Config struct {
	core.App
	core.RedisPool `mapstructure:"redis_pool"`
	olymptrade.Olymptrade
}
