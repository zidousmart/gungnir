package redis

// 哈希表命令

func (r *GRedisClient) HSet(key, field, value string) (int, error) {
	return r.Int("HSET", key, field, value)
}

func (r *GRedisClient) HSetNx(key, field, value string) (int, error) {
	return r.Int("HSETNX", key, field, value)
}

func (r *GRedisClient) HGet(key, field string) (string, error) {
	return r.String("HGET", key, field)
}

func (r *GRedisClient) HGetAll(key string) (map[string]string, error) {
	return r.StringMap("HGETALL", key)
}

func (r *GRedisClient) HExists(key string) (map[string]string, error) {
	return r.StringMap("HEXISTS", key)
}

func (r *GRedisClient) HDel(key string, fields ...interface{}) (int, error) {
	return r.Int("HDEL", key, fields)
}

func (r *GRedisClient) HKeys(key string) ([]string, error) {
	return r.Strings("HKEYS", key)
}

func (r *GRedisClient) HVals(key string) ([]string, error) {
	return r.Strings("HVALS", key)
}

func (r *GRedisClient) HLen(key string) (int, error) {
	return r.Int("HLEN", key)
}

func (r *GRedisClient) HIncrBy(key, field string, increment int) (int64, error) {
	return r.Int64("HINCRBY", key, field, increment)
}
