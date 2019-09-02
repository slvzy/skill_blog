package redis

import (
	"github.com/go-redis/redis"
	"log"
	"skill_blog/com.github.bobacsmall/pkg/setting"
)

var Client *redis.Client

func Setup() {
	Client = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password, // no password set
		DB:       0,                             // use default DB
	})
	// ping 连接是否成功
	_, err := Client.Ping().Result()
	if err != nil {
		log.Fatalf("redis.Setup, fail to client ping error : %v", err)
	}
}
