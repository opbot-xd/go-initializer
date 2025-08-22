package router

import (
	"oss.nandlabs.io/golly/rest"
)

func HealthcheckRouterHandler(server rest.Server) rest.Server {
	// Health check endpoint.
	server.Get("/v1/healthz", func(ctx rest.ServerContext) {
		ResponseJSON(ctx, 200, map[string]string{"status": "ok"})
	})

	return server
}
