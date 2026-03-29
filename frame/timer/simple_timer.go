package main

import (
	"fmt"
	"time"
)

type Task struct {
	ExecuteAt time.Time
	callback  func()
}

type SimpleTimer struct {
	tasks []*Task
}

func NewSimpleTimer() *SimpleTimer {
	return &SimpleTimer{
		tasks: []*Task{},
	}
}

func (s *SimpleTimer) AddTask(delay time.Duration, cb func()) {
	s.tasks = append(s.tasks, &Task{
		ExecuteAt: time.Now().Add(delay),
		callback:  cb,
	})
}

func (s *SimpleTimer) loop() {
	fmt.Println("begin loop")
	for {
		var earliest *Task
		var index int
		for i, v := range s.tasks {
			if earliest == nil || v.ExecuteAt.Before(earliest.ExecuteAt) {
				earliest = v
				index = i
			}
		}

		if earliest == nil {
			break
		}

		wait := time.Until(earliest.ExecuteAt)
		time.Sleep(wait)

		earliest.callback()

		s.tasks = append(s.tasks[:index], s.tasks[index+1:]...)
	}
}

func mock_simpleTimer() {

	s := NewSimpleTimer()
	// 添加定时器
	s.AddTask(2*time.Second, func() { fmt.Println("2s later") })
	s.loop()
}
