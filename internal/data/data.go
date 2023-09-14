package data

import (
	"context"
	"github.com/wastill/petclinic/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	//
	//// init mysql driver
	//_ "github.com/go-sql-driver/mysql"
)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData(client *ent.Client, rdb *redis.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	d := &Data{
		db:  client,
		rdb: rdb,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
