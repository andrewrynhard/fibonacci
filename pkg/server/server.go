package server

import (
	"log"

	"github.com/andrewrynhard/fibonacci/pkg/cache"
	sequencemodels "github.com/andrewrynhard/fibonacci/pkg/generated/server/models"
	healthzoperations "github.com/andrewrynhard/fibonacci/pkg/generated/server/restapi/operations/healthz"
	sequenceoperations "github.com/andrewrynhard/fibonacci/pkg/generated/server/restapi/operations/sequence"
	"github.com/andrewrynhard/fibonacci/pkg/healthz"
	"github.com/andrewrynhard/fibonacci/pkg/metrics"
	"github.com/andrewrynhard/fibonacci/pkg/sequence"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-redis/redis"
)

// Server is the conrete type with methods implementing the generated Swagger
// server functions.
type Server struct {
	Cache cache.Cache
}

// NewServer initializes a server with a cache.
func NewServer() *Server {
	return &Server{
		Cache: nil,
	}
}

// NewServerWithCache initializes a server with a cache.
func NewServerWithCache(cache cache.Cache) *Server {
	return &Server{
		Cache: cache,
	}
}

// GetSequence retreives a Fibonacci sequence for a value n.
func (s *Server) GetSequence(params sequenceoperations.GetSequenceParams) middleware.Responder {
	// This is checked by the math library, but we won't reach that in the case
	// of something getting inserted into the cache "accidentally".
	if params.N < 0 {
		err := sequence.NegativeNumberError{}
		return sequenceoperations.NewGetSequenceDefault(400).WithPayload(&sequencemodels.Error{Code: 400, Message: err.Error()})
	}

	// No cache layer has been configured.
	if s.Cache == nil {
		return sequence.GetSequence(params, nil)
	}

	kv, err := s.Cache.Get(params.N)
	// TODO: The Cache interfaces Get func should indicate whether or not there
	// was a cache hit/miss.
	if err == redis.Nil {
		metrics.CacheMissesCounter.Inc()
		log.Printf("cache miss: %d\n", params.N)
		return sequence.GetSequence(params, s.Cache)
	} else if err != nil {
		log.Printf("cache error: %v", err)
	}

	metrics.CacheHitsCounter.Inc()
	log.Printf("cache hit: %d\n", params.N)

	payload := &sequencemodels.Sequence{Sequence: kv.Value}

	return sequenceoperations.NewGetSequenceOK().WithPayload(payload)
}

// GetHealthz returns the health of the application.
func (s *Server) GetHealthz(params healthzoperations.GetHealthzParams) middleware.Responder {
	// TODO: This needs to be updated as the project grows.
	return healthz.GetHealthz(params)
}
