package data

import (
	"github.com/go-redis/redis/v8"
	"github.com/wastill/petclinic/internal/conf"
)

func NewRDBClient(conf *conf.Data) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
}
