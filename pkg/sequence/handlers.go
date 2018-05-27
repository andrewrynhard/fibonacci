package sequence

import (
	"log"

	"github.com/andrewrynhard/fibonacci/pkg/cache"
	"github.com/andrewrynhard/fibonacci/pkg/generated/server/models"
	"github.com/andrewrynhard/fibonacci/pkg/generated/server/restapi/operations/sequence"
	"github.com/go-openapi/runtime/middleware"
)

// GetSequence retreives a Fibonacci sequence for a value n.
func GetSequence(params sequence.GetSequenceParams, c cache.Cache) middleware.Responder {
	algo := &FastDoublingMethod{}
	n, err := Fibonacci(params.N, algo)
	if err != nil {
		return sequence.NewGetSequenceDefault(400).WithPayload(&models.Error{Code: 400, Message: err.Error()})
	}

	if c != nil {
		if err := c.Set(cache.KeyValuePair{Key: params.N, Value: n.String()}); err != nil {
			log.Printf("cache error: %v", err)
		}
	}

	nn := n.String()
	payload := &models.Sequence{N: &nn}

	return sequence.NewGetSequenceOK().WithPayload(payload)
}
