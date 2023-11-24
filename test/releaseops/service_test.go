package release_ops_test

import (
	"app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelease_Flow_Normal_1(t *testing.T) {
	input := model.ReleaseTicket{
		AppID:      "APP5",
		VersionUAT: "123",
		VersionPRD: "456",
		Job:        "ABC",
		Workflow:   "PRD",
	}
	output, err := ReleaseOPSService.Insert(&input)
	assert.ErrorIs(t, err, nil)
	err = ReleaseOPSService.TriggerBuild(output.ID)
	assert.ErrorIs(t, err, nil)
	assert.GreaterOrEqual(t, output.ID, 1)

	err = ReleaseOPSService.WorkflowSignal(output.ID, "SUCCESS")
	assert.ErrorIs(t, err, nil)
	// Approval Phase

	// Approve

	// Trigger Prod

	// Jenkins signal prod

	// Close
}
