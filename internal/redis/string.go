package redis

// 字符串命令

func (r *GRedisClient) Set(key, value string) (string, error) {
	return r.String("SET", key, value)
}

func (r *GRedisClient) Get(key string) (string, error) {
	return r.String("GET", key)
}

func (r *GRedisClient) SetEx(key, value string, seconds int) (string, error) {
	return r.String("SETEX", key, seconds, value)
}

func (r *GRedisClient) SetNx(key, value string) (int, error) {
	return r.Int("SETNX", key, value)
}

func (r *GRedisClient) PSetEx(key, value string, milliseconds int) (string, error) {
	return r.String("PSETEX", key, milliseconds, value)
}

func (r *GRedisClient) GetSet(key, value string) (string, error) {
	return r.String("GETSET", key, value)
}

func (r *GRedisClient) GetRange(key string, start, end int) (string, error) {
	return r.String("GETRANGE", key, start, end)
}

func (r *GRedisClient) Incr(key string) (int64, error) {
	return r.Int64("INCR", key)
}

func (r *GRedisClient) IncrBy(key string, increment int) (int64, error) {
	return r.Int64("INCRBY", key, increment)
}

func (r *GRedisClient) Decr(key string) (int64, error) {
	return r.Int64("DECR", key)
}

func (r *GRedisClient) DecrBy(key string, decrement int) (int64, error) {
	return r.Int64("DECRBY", key, decrement)
}

func (r *GRedisClient) Append(key, value string) (int, error) {
	return r.Int("APPEND", key, value)
}

func (r *GRedisClient) Strlen(key string) (int, error) {
	return r.Int("STRLEN", key)
}
