package container

import (
	"github.com/fgrosse/goldi"
	"github.com/rohanchauhan02/food-service/shared/config"
	"github.com/rohanchauhan02/food-service/shared/database"
)

func DefaultContainer() *goldi.Container {
	registry := goldi.NewTypeRegistry()
	conf := make(map[string]interface{})
	container := goldi.NewContainer(registry, conf)
	container.RegisterType("shared.config", config.NewImmutableConfig)
	container.RegisterType("shared.database", database.NewMysql, "@shared.config")
	return container
}
