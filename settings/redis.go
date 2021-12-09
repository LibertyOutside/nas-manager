package settings

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var rdb *redis.Client

func InitClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		log.Errorf("failed to connect redis! cause:%v", err)
		return err
	}
	return nil
}
