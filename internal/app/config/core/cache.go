package core

import (
	driverRedis "github.com/harunnryd/btrade/internal/pkg/drivers/v1.0/redis"
	"time"
)

// RedisPool ...
type RedisPool struct {
	IsEnabled          bool          `mapstructure:"is_enabled"`
	Host               string        `mapstructure:"host"`
	Port               int           `mapstructure:"port"`
	DialConnectTimeout time.Duration `mapstructure:"dial_connect_timeout"`
	ReadTimeout        time.Duration `mapstructure:"read_timeout"`
	WriteTimeout       time.Duration `mapstructure:"write_timeout"`
	MaxIdle            int           `mapstructure:"max_idle"`
	MaxActive          int           `mapstructure:"max_active"`
	IdleTimeout        time.Duration `mapstructure:"idle_timeout"`
	Wait               bool          `mapstructure:"wait"`
	MaxConnLifetime    time.Duration `mapstructure:"max_conn_lifetime"`
	Password           string        `mapstructure:"password"`
	Namespace          string        `mapstructure:"namespace"`
}

// Options ...
func (cfg *RedisPool) Options() driverRedis.Options {
	return driverRedis.Options{
		Host:               cfg.Host,
		Port:               cfg.Port,
		DialConnectTimeout: cfg.DialConnectTimeout,
		ReadTimeout:        cfg.ReadTimeout,
		WriteTimeout:       cfg.WriteTimeout,
		MaxIdle:            cfg.MaxIdle,
		MaxActive:          cfg.MaxActive,
		IdleTimeout:        cfg.IdleTimeout,
		Wait:               cfg.Wait,
		MaxConnLifetime:    cfg.MaxConnLifetime,
		Password:           cfg.Password,
		Namespace:          cfg.Namespace,
	}
}
