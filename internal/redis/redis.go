package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type GRedisClient struct {
	p *redis.Pool
}

type GRedisConfig struct {
	Address  string
	Auth     string
	Database int

	MaxIdle        int
	MaxActive      int
	IdleTimeout    time.Duration
	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

func New(config *GRedisConfig) (*GRedisClient, error) {
	pool := &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: config.IdleTimeout,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Address,
				redis.DialConnectTimeout(config.ConnectTimeout),
				redis.DialReadTimeout(config.ReadTimeout),
				redis.DialWriteTimeout(config.WriteTimeout),
			)
			if err != nil {
				return nil, err
			}
			if config.Auth != "" {
				_, err = c.Do("AUTH", config.Auth)
				if err != nil {
					c.Close()
					return nil, err
				}
			}

			if config.Database > 0 {
				_, err = c.Do("SELECT", config.Database)
				if err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, err
		},
		// TestOnBorrow: func(c redis.Conn, t time.Time) error {
		// 	if time.Since(t) < time.Minute {
		// 		return nil
		// 	}
		// 	_, err := c.Do("PING")
		// 	return err
		// },
	}

	return &GRedisClient{
		p: pool,
	}, nil
}

type GRedisStats struct {
	PoolActiveCount int
	PoolIdleCount   int
	WaitCount       int64
	WaitDuration    time.Duration
}

func (r *GRedisClient) Status() *GRedisStats {
	stats := r.p.Stats()

	return &GRedisStats{
		PoolActiveCount: stats.ActiveCount,
		PoolIdleCount:   stats.IdleCount,
		WaitCount:       stats.WaitCount,
		WaitDuration:    stats.WaitDuration,
	}
}

func (r *GRedisClient) ActiveCount() int {
	return r.p.ActiveCount()
}

func (r *GRedisClient) IdleCount() int {
	return r.p.IdleCount()
}

// 返回 int
func (r *GRedisClient) Int(cmd string, args ...interface{}) (int, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.Int(reply, e)
	if err == redis.ErrNil {
		return 0, nil
	}

	return v, err
}

// 返回 int64
func (r *GRedisClient) Int64(cmd string, args ...interface{}) (int64, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.Int64(reply, e)
	if err == redis.ErrNil {
		return 0, nil
	}

	return v, err
}

// 返回 uint64
func (r *GRedisClient) Uint64(cmd string, args ...interface{}) (uint64, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.Uint64(reply, e)
	if err == redis.ErrNil {
		return 0, nil
	}

	return v, err
}

// 返回 float64
func (r *GRedisClient) Float64(cmd string, args ...interface{}) (float64, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.Float64(reply, e)
	if err == redis.ErrNil {
		return 0, nil
	}

	return v, err
}

// 返回 string
func (r *GRedisClient) String(cmd string, args ...interface{}) (string, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.String(reply, e)
	if err == redis.ErrNil {
		return "", nil
	}

	return v, err
}

// 返回 bytes
func (r *GRedisClient) Bytes(cmd string, args ...interface{}) ([]byte, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.Bytes(reply, e)
	if err == redis.ErrNil {
		return nil, nil
	}

	return v, err
}

// 返回 bool
func (r *GRedisClient) Bool(cmd string, args ...interface{}) (bool, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.Bool(reply, e)
	if err == redis.ErrNil {
		return false, nil
	}

	return v, err
}

// 返回 []interface{}
func (r *GRedisClient) Values(cmd string, args ...interface{}) ([]interface{}, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.Values(reply, e)
	if err == redis.ErrNil {
		return nil, nil
	}

	return v, err
}

// 返回 []float64
func (r *GRedisClient) Float64s(cmd string, args ...interface{}) ([]float64, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.Float64s(reply, e)
	if err == redis.ErrNil {
		return nil, nil
	}

	return v, err
}

// 返回 []string
func (r *GRedisClient) Strings(cmd string, args ...interface{}) ([]string, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.Strings(reply, e)
	if err == redis.ErrNil {
		return nil, nil
	}

	return v, err
}

// 返回 [][]byte
func (r *GRedisClient) ByteSlices(cmd string, args ...interface{}) ([][]byte, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.ByteSlices(reply, e)
	if err == redis.ErrNil {
		return nil, nil
	}

	return v, err
}

// 返回 []int64
func (r *GRedisClient) Int64s(cmd string, args ...interface{}) ([]int64, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.Int64s(reply, e)
	if err == redis.ErrNil {
		return nil, nil
	}

	return v, err
}

// 返回 []int
func (r *GRedisClient) Ints(cmd string, args ...interface{}) ([]int, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.Ints(reply, e)
	if err == redis.ErrNil {
		return nil, nil
	}

	return v, err
}

// 返回 map[string]string
func (r *GRedisClient) StringMap(cmd string, args ...interface{}) (map[string]string, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.StringMap(reply, e)
	if err == redis.ErrNil {
		return nil, nil
	}

	return v, err
}

// 返回 map[string]int
func (r *GRedisClient) IntMap(cmd string, args ...interface{}) (map[string]int, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.IntMap(reply, e)
	if err == redis.ErrNil {
		return nil, nil
	}

	return v, err
}

// 返回 map[string]int64
func (r *GRedisClient) Int64Map(cmd string, args ...interface{}) (map[string]int64, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.Int64Map(reply, e)
	if err == redis.ErrNil {
		return nil, nil
	}

	return v, err
}

// 返回 positions
func (r *GRedisClient) Positions(cmd string, args ...interface{}) ([]*[2]float64, error) {
	conn := r.p.Get()
	defer conn.Close()

	reply, e := conn.Do(cmd, args...)
	v, err := redis.Positions(reply, e)
	if err == redis.ErrNil {
		return nil, nil
	}

	return v, err
}
