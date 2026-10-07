package rbac

import "github.com/goravel/framework/contracts/foundation"

// ServiceProvider: daftarkan SETELAH provider Gate/Auth/Orm milik framework.
type ServiceProvider struct{}

func (ServiceProvider) Register(app foundation.Application) {}

func (ServiceProvider) Boot(app foundation.Application) {
	RegisterGate()
}
