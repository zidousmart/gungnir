package redis

// 键操作命令

func (r *GRedisClient) Del(key string) (int, error) {
	return r.Int("DEL", key)
}

func (r *GRedisClient) Exists(key string) (int, error) {
	return r.Int("EXISTS", key)
}

func (r *GRedisClient) Expire(key string, seconds int) (int, error) {
	return r.Int("EXPIRE", key, seconds)
}

func (r *GRedisClient) ExpireAt(key string, timestamp int) (int, error) {
	return r.Int("EXPIREAT", key, timestamp)
}

func (r *GRedisClient) Ttl(key string) (int, error) {
	return r.Int("TTL", key)
}

func (r *GRedisClient) Persist(key string) (int, error) {
	return r.Int("PERSIST", key)
}

func (r *GRedisClient) PExpire(key string, milliseconds int64) (int, error) {
	return r.Int("PEXPIRE", key, milliseconds)
}

func (r *GRedisClient) PExpireAt(key string, timestamp int64) (int, error) {
	return r.Int("PEXPIREAT", key, timestamp)
}

func (r *GRedisClient) PTtl(key string) (int64, error) {
	return r.Int64("PTTL", key)
}
