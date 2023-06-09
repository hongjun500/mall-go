package mid

import (
	"github.com/hongjun500/mall-go/pkg/redis"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewRedisClient(t *testing.T) {
	t.Log("tests new redis client")
	assert.NotEmpty(t, client)
	t.Logf("client = %v", client)
}

func TestSet(t *testing.T) {
	t.Log("tests redis set")

	redis.SetExpiration(client, ctx, "test", "test", 5*time.Second)

	t.Logf("set success")
}

func TestGet(t *testing.T) {

}

func TestGetExpire(t *testing.T) {
	expire := redis.GetExpire(client, ctx, "test")
	t.Logf("expire = %v", expire)
}

func TestIncr(t *testing.T) {
	t.Log("tests redis increment")
	redis.SetExpiration(client, ctx, "test", 2, 5*time.Minute)
	incr := redis.Incr(client, ctx, "test")
	t.Logf("incr = %v", incr)
	t.Logf("increment success")
}

func TestHmSet(t *testing.T) {
	t.Log("tests redis hmset")
	m := make(map[string]any)
	m["test"] = "test"
	m["test1"] = "test1"
	redis.HmSetAll(client, ctx, "m", m)

	redis.HmSet(client, ctx, "m", "test", "test->修改")
	redis.HmSet(client, ctx, "m", "1", "1")

	getTest := redis.HmGet(client, ctx, "m", "test")
	assert.Contains(t, getTest, "test->修改")

	get1 := redis.HmGet(client, ctx, "m", "1")
	assert.Equal(t, "1", get1[0])

	keyTest := redis.HmHasKey(client, ctx, "m", "test")
	assert.Equal(t, true, keyTest)

	hmIncr := redis.HmIncr(client, ctx, "m", "1", 1)
	// assert.Equal(t, int64(2), hmIncr)
	t.Logf("hmIncr = %v", hmIncr)

	hmDel, i := redis.HmDel(client, ctx, "m", "test")
	assert.True(t, hmDel)
	assert.Equal(t, int64(1), i)

	hmLen := redis.HmLen(client, ctx, "m")
	assert.Equal(t, int64(2), hmLen)
	t.Logf("hmset success")

	getAll := redis.HmGetAll(client, ctx, "m")
	t.Logf("getAll = %v", getAll)
	redis.Expire(client, ctx, "m", 5*time.Second)
}

func TestHSet(t *testing.T) {
	t.Log("tests redis hset")

	add := redis.SAdd(client, ctx, "set", "test1", "test2", "test3")
	t.Logf("add = %v", add)
	members := redis.SMembers(client, ctx, "set")
	assert.Len(t, members, 3)
	t.Logf("members = %v", members)

	card := redis.SCard(client, ctx, "set")
	assert.Equal(t, int64(3), card)

	redis.SRem(client, ctx, "set", "test1")
	test1 := redis.SIsMember(client, ctx, "set", "test1")
	assert.False(t, test1)
	membersMap := redis.SMembersMap(client, ctx, "set")
	t.Logf("membersMap = %v", membersMap)

}

type Name struct {
	Name string
	Age  int
}

func TestList(t *testing.T) {
	t.Log("tests redis list")
	redis.LPush(client, ctx, "listStr", "test1", "test2", "test3")
	n1 := new(Name)
	n1.Name = "test1"
	n1.Age = 1
	n2 := new(Name)
	n2.Name = "test2"
	n2.Age = 2
	n3 := new(Name)
	n3.Name = "test3"
	n3.Age = 3
	redis.LPush(client, ctx, "listObj", n1, n2, n3)
	lRange := redis.LRange(client, ctx, "listStr", 0, -1)
	t.Logf("lRange = %v", lRange)
}
