package sequence

import (
	"log"

	"github.com/andrewrynhard/fibonacci/pkg/generated/server/models"
	"github.com/andrewrynhard/fibonacci/pkg/generated/server/restapi/operations/sequence"
	"github.com/go-openapi/runtime/middleware"
)

// GetSequence shows a sequence identified by the job's UID.
func GetSequence(params sequence.GetSequenceParams) middleware.Responder {
	algo := &FastDoublingMethod{}
	n, err := Fibonacci(params.N, algo)
	if err != nil {
		return sequence.NewGetSequenceDefault(400).WithPayload(&models.Error{Code: 400, Message: err.Error()})
	}

	log.Printf("f(%d) = %s", params.N, n.String())

	nn := n.String()
	payload := &models.Sequence{N: &nn}

	return sequence.NewGetSequenceOK().WithPayload(payload)
}
