package redis

import (
	"github.com/go-redis/redis"
	"testing"
)

func TestRedis(t *testing.T) {
	Redis.SSet("test", "pibigstar")
	t.Log(Redis.SGet("test"))
}

var lua = `
local key1 = tostring(KEYS[1])
local key2 = tostring(KEYS[2])
local args1 = tonumber(ARGV[1])
local args2 = tonumber(ARGV[2])

if key1 == "user"
then
	redis.call('SET',key1,args1)
	return 1
else
	redis.call('SET',key2,args2)
	return 2
end
return 0
`

func TestLua(t *testing.T) {
	client, err := GetRedisClient()
	if err != nil {
		t.Error(err)
	}
	script := redis.NewScript(lua)

	cmd := script.Run(client, []string{"user", "test"}, 1, 2)
	t.Log(cmd.Result())
}
