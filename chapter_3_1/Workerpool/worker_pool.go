package main

import (
	"fmt"
	"sync"
	"time"
)

// Task Definition -> for one type of task
// type Task struct {
// 	ID int
// }

// Task Definition -> for various types of tasks
type Task interface {
	Process()
}

// Email task definition
type EmailTask struct {
	Email       string
	Subject     string
	MessageBody string
}

// Image Processing task definiton
type ImgProcessingTask struct {
	ImgUrl string
}

// Way to process the tasks -> Process method (for one type)
// func (t *Task) Process() {
// 	fmt.Printf("Processing task %d\n", t.ID)

// 	//Simulate a time consuming process
// 	time.Sleep(2 * time.Second)
// }

// Way to Process Email Task
func (t *EmailTask) Process() {
	fmt.Printf("Sending email to %s\n", t.Email)

	//Simulate a time consuming process
	time.Sleep(2 * time.Second)
}

// Way to Process ImgProcessing Task
func (t *ImgProcessingTask) Process() {
	fmt.Printf("Processing the image %s\n", t.ImgUrl)

	//Simulate a time consuming process
	time.Sleep(5 * time.Second)
}

// WorkerPool Definition
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

//Functions to execute the worker pool

func (wp *WorkerPool) worker() { //Receives tasks from tasksChan and processes them
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() { //Creates goroutines and sends tasks to the channel
	//Initialise tasksChan with capacity = no of tasks
	wp.tasksChan = make(chan Task, len(wp.Tasks)) //buffered channel

	//Start workers
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	//Send tasks to the tasks channel
	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}
	close(wp.tasksChan)

	//Wait for all tasks to finish
	wp.wg.Wait()
}
