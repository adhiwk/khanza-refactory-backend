package bootstrap

import (
	"goravel/app/modules/rbac"
	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/cache"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/crypt"
	"github.com/goravel/framework/database"
	"github.com/goravel/framework/event"
	"github.com/goravel/framework/filesystem"
	"github.com/goravel/framework/hash"
	"github.com/goravel/framework/http"
	"github.com/goravel/framework/log"
	"github.com/goravel/framework/queue"
	"github.com/goravel/framework/route"
	"github.com/goravel/framework/schedule"
	"github.com/goravel/framework/session"
	"github.com/goravel/framework/validation"
	"github.com/goravel/framework/view"
	"github.com/goravel/gin"
	"github.com/goravel/mysql"
	"github.com/goravel/s3"
)

func Providers() []foundation.ServiceProvider {
	return []foundation.ServiceProvider{
		&log.ServiceProvider{},
		&cache.ServiceProvider{},
		&hash.ServiceProvider{},
		&http.ServiceProvider{},
		&session.ServiceProvider{},
		&filesystem.ServiceProvider{},
		&s3.ServiceProvider{},
		&validation.ServiceProvider{},
		&view.ServiceProvider{},
		&route.ServiceProvider{},
		&gin.ServiceProvider{},
		&database.ServiceProvider{},
		&mysql.ServiceProvider{},
		&auth.ServiceProvider{},
		&crypt.ServiceProvider{},
		&queue.ServiceProvider{},
		&event.ServiceProvider{},
		&schedule.ServiceProvider{},
		&rbac.ServiceProvider{},
	}
}
