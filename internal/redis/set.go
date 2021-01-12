package redis

// 集合命令

func (r *GRedisClient) SAdd(key string, members ...interface{}) (int, error) {
	args := make([]interface{}, 0, len(members)+1)
	args = append(args, key)
	args = append(args, members...)

	return r.Int("SADD", args...)
}

func (r *GRedisClient) SCard(key string) (int, error) {
	return r.Int("SCARD", key)
}

func (r *GRedisClient) SMembers(key string) ([]string, error) {
	return r.Strings("SMEMBERS", key)
}

func (r *GRedisClient) SIsMember(key, member string) (int, error) {
	return r.Int("SISMEMBER", key, member)
}

func (r *GRedisClient) SRem(key string, members ...interface{}) (int, error) {
	args := make([]interface{}, 0, len(members)+1)
	args = append(args, key)
	args = append(args, members...)

	return r.Int("SREM", args...)
}

func (r *GRedisClient) SPop(key string) (string, error) {
	return r.String("SPOP", key)
}

func (r *GRedisClient) SRandMember(key string, count int) ([]string, error) {
	return r.Strings("SRANDMEMBER", key, count)
}

func (r *GRedisClient) SMove(source, destination, member string) (int, error) {
	return r.Int("SMOVE", source, destination, member)
}

func (r *GRedisClient) SInter(keys ...interface{}) ([]string, error) {
	return r.Strings("SINTER", keys...)
}

func (r *GRedisClient) SInterStore(destination string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0, len(keys)+1)
	args = append(args, destination)
	args = append(args, keys...)

	return r.Int("SINTERSTORE", args...)
}

func (r *GRedisClient) SDiff(keys ...interface{}) ([]string, error) {
	return r.Strings("SDIFF", keys...)
}

func (r *GRedisClient) SDiffStore(destination string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0, len(keys)+1)
	args = append(args, destination)
	args = append(args, keys...)

	return r.Int("SDIFFSTORE", args...)
}

func (r *GRedisClient) SUnion(keys ...interface{}) ([]string, error) {
	return r.Strings("SUNION", keys...)
}

func (r *GRedisClient) SUnionStore(destination string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0, len(keys)+1)
	args = append(args, destination)
	args = append(args, keys...)

	return r.Int("SUNIONSTORE", args...)
}
