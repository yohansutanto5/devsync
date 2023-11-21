package service

import (
	"app/Integration/jenkins"
	"app/db"
	"app/model"
	"fmt"
)

// ReleaseOPSService defines the interface for managing ReleaseTickets.
type ReleaseOPSService interface {
	Insert(ReleaseTicket *model.ReleaseTicket) (ticket *model.ReleaseTicket, err error)
	GetByID(id int) (*model.ReleaseTicket, error)
	Update(data *model.ReleaseTicket) error
	JenkinsSignal(id int, signal string) error
	DeleteByID(id int) error
	TriggerBuild(id int) error
	New(FirstName, LastName string, id int) model.ReleaseTicket
	GetListTicket() ([]model.ReleaseTicket, error)
}

type ReleaseOPSServiceImpl struct {
	db *db.DataStore
}

func NewReleaseOPSService(db *db.DataStore) ReleaseOPSService {
	return &ReleaseOPSServiceImpl{db: db}
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
	err = jenkins.TriggerJenkinsWithoutParam(jenkinsURL)
	if err != nil {
		return err
	}
	// Update the ticket status into build in progress
	ticket.Status = "in progress"
	s.db.UpdateReleaseTicket(ticket)
	// return nil if there is no error
	return err
}

func (s *ReleaseOPSServiceImpl) Update(data *model.ReleaseTicket) error {
	return s.db.UpdateReleaseTicket(data)
}

func (s *ReleaseOPSServiceImpl) JenkinsSignal(id int, signal string) error {
	data, err := s.GetByID(id)
	if err != nil {
		return err
	}
	if signal == "SUCCESS" {
		data.Status = "UAT Verification"
	} else {
		data.Status = "UAT Deployment Failed"
	}
	return s.Update(data)
}

func (s *ReleaseOPSServiceImpl) Insert(ReleaseTicket *model.ReleaseTicket) (ticket *model.ReleaseTicket, err error) {
	return s.db.InsertReleaseTicket(ReleaseTicket)
}

func (s *ReleaseOPSServiceImpl) New(FirstName, LastName string, id int) model.ReleaseTicket {
	var st model.ReleaseTicket
	st.ID = id
	return st
}
