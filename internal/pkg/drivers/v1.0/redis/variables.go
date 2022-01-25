package redis

import (
	redigo "github.com/gomodule/redigo/redis"
	"sync"
)

var (
	redPoolOnce sync.Once
	redPool     *redigo.Pool
)
