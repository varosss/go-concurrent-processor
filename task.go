package main

import (
	"strconv"
	"time"
)

type Task struct {
	ID int
}

type Result struct {
	TaskID int
	Output string
	Err    error
}

func process(task Task) (string, error) {
	time.Sleep(200 * time.Millisecond)
	return "processed task #" + strconv.Itoa(task.ID), nil
}
