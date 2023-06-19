package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// 字符串部分

// SetExpiration 保存字符串并设置过期时间
func SetExpiration(client *redis.Client, ctx context.Context, key string, val any, exp time.Duration) {
	err := client.Set(ctx, key, val, exp).Err()
	if err != nil {
		log.Println("SetExpiration -> redis set error: ", err)
	}
}

// Set 保存字符串但不设置过期时间
func Set(client *redis.Client, ctx context.Context, key string, val any) {
	SetExpiration(client, ctx, key, val, 0)
}

// Get 获取字符串
func Get(client *redis.Client, ctx context.Context, key string) string {
	result, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Println("Get -> redis get error: ", err)
	}
	return result
}

// Del 批量删除字符串
func Del(client *redis.Client, ctx context.Context, keys ...string) (bool, int64) {
	result, err := client.Del(ctx, keys...).Result()
	if err != nil {
		log.Println("Del -> redis del error: ", err)
		return false, 0
	}
	return true, result
}

// Expire 设置 key 的过期时间
func Expire(client *redis.Client, ctx context.Context, key string, exp time.Duration) bool {
	err := client.Expire(ctx, key, exp).Err()
	if err != nil {
		log.Println("Expire -> redis expire error: ", err)
		return false
	}
	return err == nil
}

// GetExpire 获取 key 的过期时间
func GetExpire(client *redis.Client, ctx context.Context, key string) time.Duration {
	result, err := client.TTL(ctx, key).Result()
	if err != nil {
		log.Println("GetExpire -> redis get expire error: ", err)
		return 0
	}
	return result
}

// HasKey 判断 key 是否存在
func HasKey(client *redis.Client, ctx context.Context, key string) bool {
	result, err := client.Exists(ctx, key).Result()
	if err != nil {
		log.Println("HasKey -> redis has key error: ", err)
		return false
	}
	return result == 1
}

// Incr 自增 1
func Incr(client *redis.Client, ctx context.Context, key string) int64 {
	result, err := client.Incr(ctx, key).Result()
	if err != nil {
		log.Println("Incr -> redis incr error: ", err)
		return 0
	}
	return result
}

// IncrBy 自增指定的值
func IncrBy(client *redis.Client, ctx context.Context, key string, val int64) int64 {
	result, err := client.IncrBy(ctx, key, val).Result()
	if err != nil {
		log.Println("IncrBy -> redis incr by error: ", err)
		return 0
	}
	return result
}

// Decr 自减 1
func Decr(client *redis.Client, ctx context.Context, key string) int64 {
	result, err := client.Decr(ctx, key).Result()
	if err != nil {
		log.Println("Decr -> redis decr error: ", err)
		return 0
	}
	return result
}

// DecrBy 自减指定的值
func DecrBy(client *redis.Client, ctx context.Context, key string, val int64) int64 {
	result, err := client.DecrBy(ctx, key, val).Result()
	if err != nil {
		log.Println("DecrBy -> redis decr by error: ", err)
		return 0
	}
	return result
}

// Hash 部分

// SetNx 仅在键 key 不存在的情况下， 将键 key 的值设置为 val

// HmSet 设置哈希表 key 中的域 field 的值为 val
func HmSet(client *redis.Client, ctx context.Context, key string, field string, val any) {
	err := client.HMSet(ctx, key, field, val).Err()
	if err != nil {
		log.Println("HmSet -> redis hmset error: ", err)
	}
}

// HmSetAll 同时将多个 field-value (域-值)对设置到哈希表 key 中 (批量设置) 例如: {"name": "zhangsan", "age": 18}
func HmSetAll(client *redis.Client, ctx context.Context, key string, val map[string]any) {
	/*for field, value := range val {
		HmSet(client, ctx, key, field, value)
	}*/
	err := client.HMSet(ctx, key, val).Err()
	if err != nil {
		log.Println("HmSetAll -> redis hmset error: ", err)
	}
}

// HmGet 获取哈希表 key 中给定域 field 的值
func HmGet(client *redis.Client, ctx context.Context, key string, field string) []any {
	result, err := client.HMGet(ctx, key, field).Result()
	if err != nil {
		log.Println("HmGet -> redis hmget error: ", err)
		return nil
	}
	return result
}

// HmGetAll 获取哈希表 key 中的所有域和值
func HmGetAll(client *redis.Client, ctx context.Context, key string) map[string]string {
	result, err := client.HGetAll(ctx, key).Result()
	if err != nil {
		log.Println("HmGetAll -> redis hmgetall error: ", err)
		return nil
	}
	return result
}

// HmKeys 获取哈希表 key 中的所有域 (field)
func HmKeys(client *redis.Client, ctx context.Context, key string) []string {
	result, err := client.HKeys(ctx, key).Result()
	if err != nil {
		log.Println("HmKeys -> redis hmkeys error: ", err)
		return nil
	}
	return result
}

// HmVals 获取哈希表 key 中所有值 (value)
func HmVals(client *redis.Client, ctx context.Context, key string) []string {
	result, err := client.HVals(ctx, key).Result()
	if err != nil {
		log.Println("HmVals -> redis hmvals error: ", err)
		return nil
	}
	return result
}

// HmLen 获取哈希表 key 中域的数量
func HmLen(client *redis.Client, ctx context.Context, key string) int64 {
	result, err := client.HLen(ctx, key).Result()
	if err != nil {
		log.Println("HmLen -> redis hmlen error: ", err)
		return 0
	}
	return result
}

// HmHasKey 判断哈希表 key 中，给定域 field 是否存在
func HmHasKey(client *redis.Client, ctx context.Context, key string, field string) bool {
	result, err := client.HExists(ctx, key, field).Result()
	if err != nil {
		log.Println("HmHasKey -> redis hmhaskey error: ", err)
		return false
	}
	return result
}

// HmDel 删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略
func HmDel(client *redis.Client, ctx context.Context, key string, fields ...string) (bool, int64) {
	result, err := client.HDel(ctx, key, fields...).Result()
	if err != nil {
		log.Println("HmDel -> redis hmdel error: ", err)
		return false, 0
	}
	return true, result
}

// HmIncr 哈希表 key 中给定域 field 的值加上增量 increment
func HmIncr(client *redis.Client, ctx context.Context, key string, field string, val int64) int64 {
	result, err := client.HIncrBy(ctx, key, field, val).Result()
	if err != nil {
		log.Println("HmIncr -> redis hmincr error: ", err)
		return 0
	}
	return result
}

// HmIncrFloat 哈希表 key 中给定域 field 的值加上浮点数增量 increment
func HmIncrFloat(client *redis.Client, ctx context.Context, key string, field string, val float64) float64 {
	result, err := client.HIncrByFloat(ctx, key, field, val).Result()

	if err != nil {
		log.Println("HmIncrFloat -> redis hmincrfloat error: ", err)
		return 0
	}
	return result
}

// Set 部分

// SAdd 向集合添加一个或多个成员
func SAdd(client *redis.Client, ctx context.Context, key string, members ...any) int64 {
	result, err := client.SAdd(ctx, key, members...).Result()
	if err != nil {
		log.Println("SAdd -> redis sadd error: ", err)
		return 0
	}
	return result
}

// SMembers 获取集合中的所有成员
func SMembers(client *redis.Client, ctx context.Context, key string) []string {
	result, err := client.SMembers(ctx, key).Result()
	if err != nil {
		log.Println("SGets -> redis smembers error: ", err)
		return nil
	}
	return result
}

// SMembersMap 获取集合中的所有成员
func SMembersMap(client *redis.Client, ctx context.Context, key string) map[string]struct{} {
	result, err := client.SMembersMap(ctx, key).Result()
	if err != nil {
		log.Println("SGets -> redis smembers error: ", err)
		return nil
	}
	return result
}

// SCard 获取集合的成员数
func SCard(client *redis.Client, ctx context.Context, key string) int64 {
	result, err := client.SCard(ctx, key).Result()
	if err != nil {
		log.Println("SCard -> redis scard error: ", err)
		return 0
	}
	return result
}

// SIsMember 判断 member 元素是否是集合 key 的成员
func SIsMember(client *redis.Client, ctx context.Context, key string, member string) bool {
	result, err := client.SIsMember(ctx, key, member).Result()
	if err != nil {
		log.Println("SIsMember -> redis sismember error: ", err)
		return false
	}
	return result
}

// SRem 移除集合中一个或多个成员
func SRem(client *redis.Client, ctx context.Context, key string, members ...any) int64 {
	result, err := client.SRem(ctx, key, members...).Result()
	if err != nil {
		log.Println("SRem -> redis srem error: ", err)
		return 0
	}
	return result
}

// List 部分

// LRange 获取列表指定范围内的元素
func LRange(client *redis.Client, ctx context.Context, key string, start int64, stop int64) []string {
	result, err := client.LRange(ctx, key, start, stop).Result()
	if err != nil {
		log.Println("LRange -> redis lrange error: ", err)
		return nil
	}
	return result
}

// LLen 获取列表长度
func LLen(client *redis.Client, ctx context.Context, key string) int64 {
	result, err := client.LLen(ctx, key).Result()
	if err != nil {
		log.Println("LLen -> redis llen error: ", err)
		return 0
	}
	return result
}

// LIndex 根据索引获取列表中的元素
func LIndex(client *redis.Client, ctx context.Context, key string, index int64) string {
	result, err := client.LIndex(ctx, key, index).Result()
	if err != nil {
		log.Println("LIndex -> redis lindex error: ", err)
		return ""
	}
	return result
}

// LPush 将一个或多个值插入到列表头部
func LPush(client *redis.Client, ctx context.Context, key string, values ...any) {
	err := client.LPush(ctx, key, values...).Err()
	if err != nil {
		log.Println("Lpush -> redis lpush error: ", err)
	}
}

// LRPush 将一个或多个值插入到列表尾部
func LRPush(client *redis.Client, ctx context.Context, key string, values ...any) {
	err := client.RPush(ctx, key, values...).Err()
	if err != nil {
		log.Println("LRPush -> redis rpush error: ", err)
	}
}

// LRem 根据参数 count 的值，移除列表中与参数 value 相等的元素
func LRem(client *redis.Client, ctx context.Context, key string, count int64, value any) {
	err := client.LRem(ctx, key, count, value).Err()
	if err != nil {
		log.Println("LRem -> redis lrem error: ", err)
	}
}

// LPushX 将一个值插入到已存在的列表头部
func LPushX(client *redis.Client, ctx context.Context, key string, value any) {
	err := client.LPushX(ctx, key, value).Err()
	if err != nil {
		log.Println("LPushX -> redis lpushx error: ", err)
	}
}

// LSet  通过索引设置列表元素的值
func LSet(client *redis.Client, ctx context.Context, key string, index int64, value any) {
	err := client.LSet(ctx, key, index, value).Err()
	if err != nil {
		log.Println("LSet -> redis lset error: ", err)
	}
}
