package redis

import (
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

// NewRedis ...
func NewRedis(options Options) *redigo.Pool {
	host := fmt.Sprintf("redis://%s@%s:%d", options.Password, options.Host, options.Port)

	redPoolOnce.Do(func() {
		redPool = &redigo.Pool{
			Dial: func() (redigo.Conn, error) {
				c, err := redigo.DialURL(host)
				if err != nil {
					return nil, err
				}

				if options.Password != "" {
					if _, err := c.Do("AUTH", options.Password); err != nil {
						c.Close()
						return nil, err
					}
				}

				if _, err := c.Do("SELECT", options.Namespace); err != nil {
					c.Close()
					return nil, err
				}
				return c, nil
			},
			TestOnBorrow: func(c redigo.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
			MaxIdle:         options.MaxIdle,
			MaxActive:       options.MaxActive,
			IdleTimeout:     options.IdleTimeout * time.Second,
			Wait:            options.Wait,
			MaxConnLifetime: options.MaxConnLifetime * time.Second,
		}

		conn := redPool.Get()
		defer conn.Close()

		_, err := conn.Do("PING")
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	return redPool
}
