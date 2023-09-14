package data

import (
	"github.com/go-redis/redis/v8"
	"github.com/wastill/petclinic/internal/conf"
)

func NewRDBClient(conf *conf.Bootstrap) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         conf.Data.Redis.Addr,
		Password:     conf.Data.Redis.Password,
		DB:           int(conf.Data.Redis.Db),
		DialTimeout:  conf.Data.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Data.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Data.Redis.ReadTimeout.AsDuration(),
	})
}
