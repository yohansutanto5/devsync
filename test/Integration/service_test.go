package integration_test

import (
	"app/Integration/jenkins"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriggerJenkins(t *testing.T) {
	err := jenkins.TriggerJenkinsWithoutParam()
	assert.ErrorIs(t, nil, err)
}
