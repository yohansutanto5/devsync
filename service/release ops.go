package service

import (
	"app/Integration"
	"app/constanta"
	"app/db"
	"app/model"
	"fmt"
)

// ReleaseOPSService defines the interface for managing ReleaseTickets.
type ReleaseOPSService interface {
	Insert(ReleaseTicket *model.ReleaseTicket) (ticket *model.ReleaseTicket, err error)
	GetByID(id int) (*model.ReleaseTicket, error)
	Update(data *model.ReleaseTicket) error
	WorkflowSignal(id int, signal string) error
	DeleteByID(id int) error
	TriggerBuild(id int) error
	New(FirstName, LastName string, id int) model.ReleaseTicket
	GetListTicket() ([]model.ReleaseTicket, error)
}

type ReleaseOPSServiceImpl struct {
	db       *db.DataStore
	external *Integration.ExternalService
}

func NewReleaseOPSService(db *db.DataStore, external *Integration.ExternalService) ReleaseOPSService {
	return &ReleaseOPSServiceImpl{db: db, external: external}
}

// Function Implementation
func (s *ReleaseOPSServiceImpl) GetByID(id int) (*model.ReleaseTicket, error) {
	// Implementation for fetching a ReleaseTicket by ID from the database
	return s.db.GetReleaseTicketByID(id)
}

func (s ReleaseOPSServiceImpl) GetListTicket() ([]model.ReleaseTicket, error) {
	return s.db.GetListReleaseTicket()
}

func (s *ReleaseOPSServiceImpl) DeleteByID(id int) error {
	return s.db.DeleteReleaseTicketByID(id)
}

func (s *ReleaseOPSServiceImpl) TriggerBuild(id int) error {
	// Retrieve the Jenkins Data from Ticket Data
	fmt.Println(id)
	ticket, err := s.GetByID(id)
	if err != nil {
		return err
	}
	jenkinsURL := "https://staging-jenkins.nexcloud.id/job/devsync/job/Credential/build"

	// Trigger jenkins >> This process is ASYNC
	err = s.external.TriggerJenkinsWithoutParam(jenkinsURL)
	if err != nil {
		return err
	}
	// Update the ticket status into build in progress
	if ticket.Status == constanta.PRDReady {
		ticket.Status = constanta.PRDDeploy
	} else if ticket.Status == constanta.UATReady {
		ticket.Status = constanta.UATDeploy
	} else {
		// make an error state that Status is not match
		return err
	}
	return s.db.UpdateReleaseTicket(ticket)
}

func (s *ReleaseOPSServiceImpl) Update(data *model.ReleaseTicket) error {
	return s.db.UpdateReleaseTicket(data)
}

// WorkflowSignal handles the workflow of the CI/CD pipeline, signaling different stages.
// Flow Chart:
//   UAT Deployment Ready --> UAT Deployment In Progress --> (1) / (2)
//   (1) UAT Deployment Verified *SUCCESS SIGNAL* -->  (3) / (4)
//   (2) Deployment Failed *FAILED SIGNAL* --> (5) / (6)

// CheckBuildStatus --> TriggerTests --> RunTests --> CheckTestStatus -->
// TriggerDeployment --> DeployCode --> CheckDeploymentStatus --> End
func (s *ReleaseOPSServiceImpl) WorkflowSignal(id int, signal string) error {
	data, err := s.GetByID(id)
	if err != nil {
		return err
	}
	//
	if signal == "SUCCESS" {
		data.Status = constanta.UATVerify
	} else if signal == "FAILED" {
		data.Status = constanta.Failed
	} else if signal == "APPROVED" {
		data.Status = constanta.PRDReady
	} else if signal == "REJECTED" {
		data.Status = constanta.Rejected
	} else if data.Status == constanta.UATVerify && signal == "VERIFIED" {
		data.Status = constanta.ApprovalPending
	} else if data.Status == constanta.PRDVerify && signal == "VERIFIED" {
		data.Status = constanta.Closed
	} else {
		// Create an error that status is missmatch
		return err
	}
	return s.Update(data)
}

func (s *ReleaseOPSServiceImpl) Insert(ReleaseTicket *model.ReleaseTicket) (ticket *model.ReleaseTicket, err error) {
	ReleaseTicket.Status = constanta.UATReady
	return s.db.InsertReleaseTicket(ReleaseTicket)
}

func (s *ReleaseOPSServiceImpl) New(FirstName, LastName string, id int) model.ReleaseTicket {
	var st model.ReleaseTicket
	st.ID = id
	return st
}
