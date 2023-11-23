package Integration

import (
	"app/Integration/jenkins"
	"app/cmd/config"
)

// ExternalService defines the interface for managing Externals

type ExternalService struct {
	config *config.Configuration
}

func NewExternalService(config *config.Configuration) *ExternalService {
	return &ExternalService{config: config}
}

// Function Implementation

func (s ExternalService) TriggerJenkinsWithoutParam(url string) error {
	return jenkins.TriggerJenkinsWithoutParam(url)
}
