package bootstrap

import (
	"github.com/goravel/framework/contracts/database/seeder"
	contractsfoundation "github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/foundation"

	"goravel/config"
	"goravel/database/seeders"
	"goravel/routes"
)

func Boot() contractsfoundation.Application {
	return foundation.Setup().
		WithMigrations(Migrations).
		WithSeeders(Seeders).
		WithRouting(func() {
			routes.Web()
			routes.Api()
		}).
		WithProviders(Providers).
		WithConfig(config.Boot).
		Create()
}

func Seeders() []seeder.Seeder {
	return []seeder.Seeder{
		&seeders.RbacSeeder{},
	}
}
