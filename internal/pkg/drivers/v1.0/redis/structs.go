package redis

import "time"

// Options ...
type Options struct {
	Host               string
	Port               int
	DialConnectTimeout time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxIdle            int
	MaxActive          int
	IdleTimeout        time.Duration
	Wait               bool
	MaxConnLifetime    time.Duration
	Password           string
	Namespace          string
}
