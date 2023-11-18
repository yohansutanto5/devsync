package testscheduler

import (
	"app/scheduler"
	"testing"
)

// func TestMain(m *testing.M) {
// 	m.Run()
// }

func TestSetup(t *testing.T) {
	scheduler.Setup()
}
