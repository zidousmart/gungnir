package redis

// 列表命令

func (r *GRedisClient) LPush(key string, values ...interface{}) (int, error) {
	args := make([]interface{}, 0, len(values)+1)
	args = append(args, key)
	args = append(args, values...)

	return r.Int("LPUSH", args...)
}

func (r *GRedisClient) RPush(key string, values ...interface{}) (int, error) {
	args := make([]interface{}, 0, len(values)+1)
	args = append(args, key)
	args = append(args, values...)

	return r.Int("RPUSH", args...)
}

func (r *GRedisClient) LPushX(key string, value string) (int, error) {
	return r.Int("LPUSHX", key, value)
}

func (r *GRedisClient) RPushX(key string, value string) (int, error) {
	return r.Int("RPUSHX", key, value)
}

func (r *GRedisClient) LPop(key string) (string, error) {
	return r.String("LPOP", key)
}

func (r *GRedisClient) RPop(key string) (string, error) {
	return r.String("RPOP", key)
}

func (r *GRedisClient) RPopLPush(source, destination string) (string, error) {
	return r.String("RPOPLPUSH", source, destination)
}

func (r *GRedisClient) LLen(key string) (int, error) {
	return r.Int("LLEN", key)
}

func (r *GRedisClient) LIndex(key string, index int) (string, error) {
	return r.String("LINDEX", key, index)
}

func (r *GRedisClient) LRange(key string, start, stop int) ([]string, error) {
	return r.Strings("LRANGE", key, start, stop)
}

func (r *GRedisClient) LRem(key string, count int, value string) (int, error) {
	return r.Int("LREM", key, count, value)
}

func (r *GRedisClient) LSet(key string, index int, value string) (bool, error) {
	reply, err := r.String("LSET", key, index, value)
	if err != nil {
		return false, err
	}

	return reply == "OK", nil
}
