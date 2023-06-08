package mid

import (
	"github.com/hongjun500/mall-go/pkg/redis"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	redis.NewRedisClient()
	m.Run()
}

func TestNewRedisClient(t *testing.T) {
	t.Log("tests new redis client")
	client := redis.NewRedisClient()
	t.Logf("client = %v", client)
}

func TestSet(t *testing.T) {
	t.Log("tests redis set")

	err := redis.SetExpiration("test", "test", 5*time.Second)
	if err != nil {
		t.Logf("set error, err = %v", err)
		return
	}
	t.Logf("set success")
}
