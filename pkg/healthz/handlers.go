package healthz

import (
	"github.com/andrewrynhard/fibonacci/pkg/generated/server/restapi/operations/healthz"
	"github.com/go-openapi/runtime/middleware"
)

// GetHealthz returns the health of the application.
func GetHealthz(params healthz.GetHealthzParams) middleware.Responder {
	// TODO: This needs to be updated as the project grows.
	return healthz.NewGetHealthzOK()
}
