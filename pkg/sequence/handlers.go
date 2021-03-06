package sequence

import (
	"fmt"
	"log"
	"time"

	"github.com/andrewrynhard/fibonacci/pkg/cache"
	"github.com/andrewrynhard/fibonacci/pkg/generated/server/models"
	"github.com/andrewrynhard/fibonacci/pkg/generated/server/restapi/operations/sequence"
	"github.com/andrewrynhard/fibonacci/pkg/metrics"
	"github.com/go-openapi/runtime/middleware"
)

// GetSequence retreives a Fibonacci sequence for a value n.
func GetSequence(params sequence.GetSequenceParams, c cache.Cache) middleware.Responder {
	payload := &models.Sequence{}
	algo := &FastDoublingMethod{}
	start := time.Now()
	for i := int64(0); i < (params.N); i++ {
		n, err := Fibonacci(i, algo)
		if err != nil {
			return sequence.NewGetSequenceDefault(400).WithPayload(&models.Error{Code: 400, Message: err.Error()})
		}
		payload.Sequence = append(payload.Sequence, n.String())
	}
	duration := time.Since(start)
	metrics.DurationHistogram.WithLabelValues(fmt.Sprintf("%d", 200)).Observe(duration.Seconds())

	if c != nil {
		if err := c.Set(cache.KeyValuePair{Key: params.N, Value: payload.Sequence}); err != nil {
			log.Printf("cache error: %v", err)
		}
	}

	return sequence.NewGetSequenceOK().WithPayload(payload)
}
