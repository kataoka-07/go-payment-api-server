package generator

import (
	"fmt"
	"go-payment-api-server/pkg/logger"

	"github.com/samber/lo"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func GenerateModel(cfg GenConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FTokyo",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Error("failed to connect DB", "err", err)
		lo.Must0(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "internal/infrastructure/query",
		ModelPkgPath:      "internal/domain/model",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
