package common

import (
	"testing"
	"time"

	"github.com/hongjun500/mall-go/pkg/redis"
	"github.com/stretchr/testify/assert"
)

func TestNewRedisClient(t *testing.T) {
	t.Log("tests new redis client")
	assert.NotEmpty(t, redisCli)
	t.Logf("redisCli = %v", redisCli)
}

func TestSet(t *testing.T) {
	t.Log("tests redis set")

	redis.SetExpiration(redisCli, ctx, "test", "test", 5*time.Second)

	t.Logf("set success")
}

func TestGet(t *testing.T) {

}

func TestGetExpire(t *testing.T) {
	expire := redis.GetExpire(redisCli, ctx, "test")
	t.Logf("expire = %v", expire)
}

func TestIncr(t *testing.T) {
	t.Log("tests redis increment")
	redis.SetExpiration(redisCli, ctx, "test", 2, 5*time.Minute)
	incr := redis.Incr(redisCli, ctx, "test")
	t.Logf("incr = %v", incr)
	t.Logf("increment success")
}

func TestHmSet(t *testing.T) {
	t.Log("tests redis hmset")
	m := make(map[string]any)
	m["test"] = "test"
	m["test1"] = "test1"
	redis.HmSetAll(redisCli, ctx, "m", m)

	redis.HmSet(redisCli, ctx, "m", "test", "test->修改")
	redis.HmSet(redisCli, ctx, "m", "1", "1")

	getTest := redis.HmGet(redisCli, ctx, "m", "test")
	assert.Contains(t, getTest, "test->修改")

	get1 := redis.HmGet(redisCli, ctx, "m", "1")
	assert.Equal(t, "1", get1[0])

	keyTest := redis.HmHasKey(redisCli, ctx, "m", "test")
	assert.Equal(t, true, keyTest)

	hmIncr := redis.HmIncr(redisCli, ctx, "m", "1", 1)
	// assert.Equal(t, int64(2), hmIncr)
	t.Logf("hmIncr = %v", hmIncr)

	hmDel, i := redis.HmDel(redisCli, ctx, "m", "test")
	assert.True(t, hmDel)
	assert.Equal(t, int64(1), i)

	hmLen := redis.HmLen(redisCli, ctx, "m")
	assert.Equal(t, int64(2), hmLen)
	t.Logf("hmset success")

	getAll := redis.HmGetAll(redisCli, ctx, "m")
	t.Logf("getAll = %v", getAll)
	redis.Expire(redisCli, ctx, "m", 5*time.Second)
}

func TestHSet(t *testing.T) {
	t.Log("tests redis hset")

	add := redis.SAdd(redisCli, ctx, "set", "test1", "test2", "test3")
	t.Logf("add = %v", add)
	members := redis.SMembers(redisCli, ctx, "set")
	assert.Len(t, members, 3)
	t.Logf("members = %v", members)

	card := redis.SCard(redisCli, ctx, "set")
	assert.Equal(t, int64(3), card)

	redis.SRem(redisCli, ctx, "set", "test1")
	test1 := redis.SIsMember(redisCli, ctx, "set", "test1")
	assert.False(t, test1)
	membersMap := redis.SMembersMap(redisCli, ctx, "set")
	t.Logf("membersMap = %v", membersMap)

}

type Name struct {
	Name string
	Age  int
}

func TestList(t *testing.T) {
	t.Log("tests redis list")
	redis.LPush(redisCli, ctx, "listStr", "test1", "test2", "test3")
	n1 := new(Name)
	n1.Name = "test1"
	n1.Age = 1
	n2 := new(Name)
	n2.Name = "test2"
	n2.Age = 2
	n3 := new(Name)
	n3.Name = "test3"
	n3.Age = 3
	redis.LPush(redisCli, ctx, "listObj", n1, n2, n3)
	lRange := redis.LRange(redisCli, ctx, "listStr", 0, -1)
	t.Logf("lRange = %v", lRange)
}
