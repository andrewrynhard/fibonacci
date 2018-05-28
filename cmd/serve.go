// Copyright © 2018 Andrew Rynhard <andrew@andrewrynhard.com>

package cmd

import (
	"log"
	"os"

	"github.com/andrewrynhard/fibonacci/pkg/cache"
	"github.com/andrewrynhard/fibonacci/pkg/cache/redis"
	"github.com/andrewrynhard/fibonacci/pkg/generated/server/restapi"
	"github.com/andrewrynhard/fibonacci/pkg/generated/server/restapi/operations"
	"github.com/andrewrynhard/fibonacci/pkg/generated/server/restapi/operations/healthz"
	"github.com/andrewrynhard/fibonacci/pkg/generated/server/restapi/operations/sequence"
	"github.com/andrewrynhard/fibonacci/pkg/metrics"
	"github.com/andrewrynhard/fibonacci/pkg/server"
	"github.com/andrewrynhard/fibonacci/pkg/ui"
	loads "github.com/go-openapi/loads"
	"github.com/spf13/cobra"
)

var (
	useCache    bool
	apiPort     int
	metrcisPort int
	uiPort      int
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  ``,
}

// apiCmd represents the serve api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
		if err != nil {
			log.Fatalln(err)
		}

		go func() {
			if err := metrics.ServeMetrics(metrcisPort); err != nil {
				log.Printf("failed to start metrics: %v", err)
			}
		}()

		var fibserver *server.Server
		if useCache {
			var c cache.Cache
			// TODO: What happens if we add another cache type? This needs a
			// better design.
			if endpoint, ok := os.LookupEnv("REDIS_ENDPOINT"); ok {
				log.Printf("using cache at %s", endpoint)
				c = redis.NewRedisCache(endpoint)
			}
			fibserver = server.NewServerWithCache(c)
		} else {
			fibserver = server.NewServer()
		}

		api := operations.NewFibonacciAPI(swaggerSpec)
		api.SequenceGetSequenceHandler = sequence.GetSequenceHandlerFunc(fibserver.GetSequence)
		api.HealthzGetHealthzHandler = healthz.GetHealthzHandlerFunc(fibserver.GetHealthz)

		server := restapi.NewServer(api)
		// nolint: errcheck
		defer server.Shutdown()

		server.ConfigureAPI()
		server.Port = apiPort

		if err := server.Serve(); err != nil {
			log.Fatalln(err)
		}
	},
}

// uiCmd represents the serve ui command
var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := ui.Serve(uiPort); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	apiCmd.Flags().BoolVar(&useCache, "with-cache", false, "Use a caching layer")
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().IntVar(&apiPort, "api-port", 80, "Port to listen on (default: 80)")
	serveCmd.PersistentFlags().IntVar(&metrcisPort, "metrics-port", 82, "Port to listen on (default: 82)")
	serveCmd.PersistentFlags().IntVar(&uiPort, "ui-port", 81, "Port to listen on (default: 81)")
	serveCmd.AddCommand(apiCmd, uiCmd)
}
