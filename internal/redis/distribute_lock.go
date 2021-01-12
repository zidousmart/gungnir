package redis

import (
	"errors"

	redigo "github.com/gomodule/redigo/redis"
)

// 分布式锁
// TODO 延长锁

// 加锁
func (r *GRedisClient) Lock(expire int, key, value string) (bool, error) {
	if expire > 0 {
		reply, err := r.String("SET", key, value, "EX", expire, "NX")
		if err != nil {
			return false, err
		}

		return reply == "OK", nil
	}

	return false, errors.New("expire second not set")
}

// lua脚本，用来释放分布式锁
var luaUnLock = "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end"

// 解锁
func (r *GRedisClient) UnLock(key, value string) (bool, error) {
	conn := r.p.Get()
	defer conn.Close()

	lua := redigo.NewScript(1, luaUnLock)
	reply, err := redigo.Int(lua.Do(conn, key, value))
	if err != nil {
		return false, err
	}

	if reply > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
