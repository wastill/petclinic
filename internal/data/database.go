package data

//import (
//	"context"
//	"entgo.io/ent"
//	"entgo.io/ent/dialect/sql/schema"
//	"github.com/go-kratos/kratos/v2/config"
//	"log"
//	pkgent "github.com/wastill/petclinic/internal/data/ent"
//)
//
//func NewEnt(cfg config.Config) *ent.Client {
//	dbDrvier, err := pkgent.InitDatabaseDriver(cfg.GetString("database.wms.type"), cfg.GetString("database.wms.connect_uri"))
//	if err != nil {
//		log.Fatal("database init err: " + err.Error())
//	}
//	// init ent
//	db := ent.NewClient(ent.Driver(dbDrvier))
//	if !test_flag.AutoMigration {
//		return db
//	}
//	// Run the auto migration tool.
//	if err := db.Schema.Create(
//		context.Background(),
//		schema.WithAtlas(true),
//		migrate.WithDropIndex(true),
//		migrate.WithDropColumn(true),
//	); err != nil {
//		log.Fatal("failed creating schema resources: " + err.Error())
//	}
//
//	return db
//}
