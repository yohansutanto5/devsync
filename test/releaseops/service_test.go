package release_ops_test

import (
	"app/constanta"
	"app/model"
	"app/pkg/log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelease_Flow_Normal_1(t *testing.T) {
	input := model.ReleaseTicket{
		AppID:      "APP5",
		VersionUAT: "125",
		VersionPRD: "456",
		Job:        "ABC",
		Workflow:   "PRD",
	}
	// Input Initial ticket
	output, err := ReleaseOPSService.Insert(&input)
	assert.ErrorIs(t, err, nil)
	assert.GreaterOrEqual(t, output.ID, 1)

	// Trigger Deploy UAT
	err = ReleaseOPSService.TriggerDeploy(output.ID)
	assert.ErrorIs(t, err, nil)

	// Signal UAT Deploy Success
	err = ReleaseOPSService.WorkflowSignal(output.ID, "SUCCESS")
	assert.ErrorIs(t, err, nil)

	// Signal UAT Verified
	err = ReleaseOPSService.WorkflowSignal(output.ID, "VERIFIED")
	assert.ErrorIs(t, err, nil)

	// Signal Approve Prod Deploymen
	err = ReleaseOPSService.WorkflowSignal(output.ID, "APPROVED")
	assert.ErrorIs(t, err, nil)

	// Signal Trigger Prod
	err = ReleaseOPSService.TriggerDeploy(output.ID)
	assert.ErrorIs(t, err, nil)

	// Jenkins signal prod deploy Success
	err = ReleaseOPSService.WorkflowSignal(output.ID, "SUCCESS")
	assert.ErrorIs(t, err, nil)

	// Close Ticket after verify prod Deployment
	err = ReleaseOPSService.WorkflowSignal(output.ID, "VERIFIED")
	assert.ErrorIs(t, err, nil)

	// Get the Ticket Details
	Newinput, _ := ReleaseOPSService.GetByID(output.ID)
	assert.Equal(t, constanta.Closed, Newinput.Status)
	log.Warning(0, "Debug Data Flow", Newinput)
}

func TestRelease_Flow_Rejected_1(t *testing.T) {
	input := model.ReleaseTicket{
		AppID:      "APP5",
		VersionUAT: "125",
		VersionPRD: "456",
		Job:        "ABC",
		Workflow:   "PRD",
	}
	// Input Initial ticket
	output, err := ReleaseOPSService.Insert(&input)
	assert.ErrorIs(t, err, nil)
	assert.GreaterOrEqual(t, output.ID, 1)

	// Trigger Deploy UAT
	err = ReleaseOPSService.TriggerDeploy(output.ID)
	assert.ErrorIs(t, err, nil)

	// Signal UAT Deploy Success
	err = ReleaseOPSService.WorkflowSignal(output.ID, "SUCCESS")
	assert.ErrorIs(t, err, nil)

	// Signal UAT Verified
	err = ReleaseOPSService.WorkflowSignal(output.ID, "VERIFIED")
	assert.ErrorIs(t, err, nil)

	// Signal Approve Prod Deploymen
	err = ReleaseOPSService.WorkflowSignal(output.ID, "REJECTED")
	assert.ErrorIs(t, err, nil)

	// Signal Trigger Prod
	err = ReleaseOPSService.TriggerDeploy(output.ID)
	assert.NotNil(t, err, nil)

	// Get the Ticket Details
	Newinput, _ := ReleaseOPSService.GetByID(output.ID)
	assert.Equal(t, constanta.Rejected, Newinput.Status)
}

func TestRelease_Flow_Deploy_Failed_1(t *testing.T) {
	input := model.ReleaseTicket{
		AppID:      "APP5",
		VersionUAT: "125",
		VersionPRD: "456",
		Job:        "ABC",
		Workflow:   "PRD",
	}
	// Input Initial ticket
	output, err := ReleaseOPSService.Insert(&input)
	assert.ErrorIs(t, err, nil)
	assert.GreaterOrEqual(t, output.ID, 1)

	// Trigger Deploy UAT
	err = ReleaseOPSService.TriggerDeploy(output.ID)
	assert.ErrorIs(t, err, nil)

	// Signal UAT Deploy Success
	err = ReleaseOPSService.WorkflowSignal(output.ID, "FAILED")
	assert.ErrorIs(t, err, nil)

	// Signal UAT Verified should be failed
	err = ReleaseOPSService.WorkflowSignal(output.ID, "VERIFIED")
	assert.NotNil(t, err, nil)

	// Get the Ticket Details
	Newinput, _ := ReleaseOPSService.GetByID(output.ID)
	assert.Equal(t, constanta.Failed, Newinput.Status)
}
