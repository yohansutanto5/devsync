package scheduler

import (
	"app/pkg/log"
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func Setup() {
	gocron.Every(2).Seconds().Do(task1)
	<-gocron.Start()
	select {}
}

func task1() {
	fmt.Println("Task 1 Executed")
	log.Debug(1, "task 1 executed", nil)
}
