package data

import (
	"context"
	config "github.com/wastill/petclinic/internal/conf"
	"github.com/wastill/petclinic/internal/data/ent"
	"github.com/wastill/petclinic/internal/data/ent/migrate"
	pkgent "github.com/wastill/petclinic/pkg/ent"
	"log"
)

func NewEnt(cfg *config.Data) *ent.Client {
	dbDrvier, err := pkgent.InitDatabaseDriver(cfg.Database.Driver, cfg.Database.Source)
	if err != nil {
		log.Fatal("database init err: " + err.Error())
	}
	// init ent
	db := ent.NewClient(ent.Driver(dbDrvier))
	if cfg.Database.AutoMigrate {
		// Run the auto migration tool.
		if err = db.Schema.Create(
			context.Background(),
			// 已默认使用Atlas,不用代码声明
			// schema.WithAtlas(true),
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
		); err != nil {
			log.Fatal("failed creating schema resources: " + err.Error())
		}
	}

	return db
}
