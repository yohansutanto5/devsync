package integration_test

import (
	"app/Integration/jenkins"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriggerJenkins(t *testing.T) {
	jenkinsURL := "https://staging-jenkins.nexcloud.id/job/devsync/job/Credential/build"
	err := jenkins.TriggerJenkinsWithoutParam(jenkinsURL)
	assert.ErrorIs(t, nil, err)
}
