package redis

import (
	"errors"
)

// 有序集合命令

func (r *GRedisClient) ZAdd(key string, scores []int, members []string) (int, error) {
	if len(scores) != len(members) {
		return 0, errors.New("param error")
	}

	args := make([]interface{}, 0, len(scores)*2+1)
	args = append(args, key)
	for i, score := range scores {
		args = append(args, score)
		args = append(args, members[i])
	}

	return r.Int("ZADD", args...)
}

func (r *GRedisClient) ZCard(key string) (int, error) {
	return r.Int("ZCARD", key)
}

func (r *GRedisClient) ZCount(key string, min, max int) (int, error) {
	return r.Int("ZCOUNT", key, min, max)
}

func (r *GRedisClient) ZIncrBy(key, member string, increment int) (int64, error) {
	return r.Int64("ZINCRBY", key, increment, member)
}

func (r *GRedisClient) ZRange(key string, start, stop int, withScores bool) ([]string, error) {
	var args []interface{}
	args = append(args, key)
	args = append(args, start)
	args = append(args, stop)
	if withScores {
		args = append(args, "WITHSCORES")
	}

	return r.Strings("ZRANGE", args...)
}

func (r *GRedisClient) ZRangeByScore(key string, min, max int, withScores bool) ([]string, error) {
	var args []interface{}
	args = append(args, key)
	args = append(args, min)
	args = append(args, max)
	if withScores {
		args = append(args, "WITHSCORES")
	}

	return r.Strings("ZRANGEBYSCORE", args...)
}

func (r *GRedisClient) ZRank(key, member string) (int, error) {
	return r.Int("ZRANK", key, member)
}

func (r *GRedisClient) ZRem(key string, members ...interface{}) (int, error) {
	args := make([]interface{}, 0, len(members)+1)
	args = append(args, key)
	args = append(args, members...)

	return r.Int("ZREM", args...)
}

func (r *GRedisClient) ZRemRangeByRank(key string, start, stop int) (int, error) {
	return r.Int("ZREMRANGEBYRANK", key, start, stop)
}

func (r *GRedisClient) ZRemRangeByScore(key string, min, max int) (int, error) {
	return r.Int("ZREMRANGEBYSCORE", key, min, max)
}

func (r *GRedisClient) ZRevRange(key string, start, stop int, withScores bool) ([]string, error) {
	var args []interface{}
	args = append(args, key)
	args = append(args, start)
	args = append(args, stop)
	if withScores {
		args = append(args, "WITHSCORES")
	}

	return r.Strings("ZREVRANGE", args...)
}

func (r *GRedisClient) ZRevRangeByScore(key string, max, min int, withScores bool) ([]string, error) {
	var args []interface{}
	args = append(args, key)
	args = append(args, max)
	args = append(args, min)
	if withScores {
		args = append(args, "WITHSCORES")
	}

	return r.Strings("ZREVRANGEBYSCORE", args...)
}

func (r *GRedisClient) ZRevRank(key, member string) (int, error) {
	return r.Int("ZREVRANK", key, member)
}

func (r *GRedisClient) ZScore(key, member string) (int, error) {
	return r.Int("ZSCORE", key, member)
}
