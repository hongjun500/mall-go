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
