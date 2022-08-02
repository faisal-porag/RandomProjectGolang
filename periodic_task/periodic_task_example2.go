package main

import (
	"fmt"
	"time"

	"github.com/jtaczanowski/go-scheduler"
)

// RUN THIS COMMAND BEFORE EXECUTE THIS PROGRAM
// go get github.com/jtaczanowski/go-scheduler

func main() {
	// run function exampleTask() every 10s
	scheduler.RunTaskAtInterval(exampleTask, time.Second*10, time.Second*0)

	// run function exampleTask() every 10s, but delay on start by 3 second
	scheduler.RunTaskAtInterval(exampleTask, time.Second*10, time.Second*3)

	// run function exampleTaskWithArguments("example task with arguments") every 10s
	scheduler.RunTaskAtInterval(
		func() { exampleTaskWithArguments("example task with arguments") },
		time.Second*10,
		time.Second*0,
	)
}

func exampleTask() {
	fmt.Print("example")
}

func exampleTaskWithArguments(str string) {
	fmt.Print(str)
}
