package rbac

import (
	"context"

	"github.com/goravel/framework/auth/access"
	contractsaccess "github.com/goravel/framework/contracts/auth/access"
	"github.com/goravel/framework/facades"
)

// RegisterGate menghubungkan RBAC ke Gate Goravel.
// Setelah ini: facades.Gate().WithContext(ctx).Allows("users.create", nil)
// bernilai true bila user punya permission tsb (atau super-admin).
// Return nil = "tidak berpendapat", sehingga Gate::Define milik user login tetap jalan.
func RegisterGate() {
	facades.Gate().Before(func(ctx context.Context, ability string, arguments map[string]any) contractsaccess.Response {
		uid, ok := userIDFromContext(ctx)
		if !ok {
			return nil
		}
		if Default().Can(uid, ability) {
			return access.NewAllowResponse()
		}
		return nil
	})
}
