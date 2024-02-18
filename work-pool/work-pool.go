package main

import (
	"fmt"
	"sync"
	"time"
)

type Task interface {
	Process()
}

type EmailTask struct {
	Email   string
	Subject string
	Message string
}

func (t *EmailTask) Process() {
	fmt.Printf("Sending email to %s\n", t.Email)
	time.Sleep(5 * time.Second)
}

type ImageTask struct {
	ImageUrl string
}

func (t *ImageTask) Process() {
	fmt.Printf("Processing the image %s\n", t.ImageUrl)
	time.Sleep(10 * time.Second)
}

type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

// func (t *Task) Process() {
// 	fmt.Printf("Processing task %d\n", t.ID)

// 	time.Sleep(3 * time.Second)
// }

func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	wp.tasksChan = make(chan Task, len(wp.Tasks))
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	wp.wg.Add(len(wp.Tasks))

	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}

	close(wp.tasksChan)

	wp.wg.Wait()
}
