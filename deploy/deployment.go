package main

import (
	"io"
)

// Deployment is the deployment object
type Deployment struct {
	Environment string
	Service     string
	Manifests   map[string]io.Reader
}

// NewDeployment creates a Deployment type
func NewDeployment(environment string, service string) (Deployment, error) {
	return Deployment{}, nil
}

// Apply apply manifests in kubernetes
func (d *Deployment) Apply() error {
	return nil
}
