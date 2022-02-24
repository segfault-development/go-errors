package errors

import (
	"errors"
	"fmt"
)

type integrationType string

const (
	// Used for communication errors
	IntegrationTypeNetwork integrationType = "Network"
	// Used for blockchain errors like Hyperledger
	IntegrationTypeBlockchain integrationType = "Blockchain"
)

var (
	errIntegration = errors.New("Integration")
)

func NewIntegration(integration integrationType, msg string) error {
	return fmt.Errorf("%s: %s: %w", integration, msg, errIntegration)
}

func IsIntegration(err error) bool {
	return errors.Is(err, errIntegration)
}
