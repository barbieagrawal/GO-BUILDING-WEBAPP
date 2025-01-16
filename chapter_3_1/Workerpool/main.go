package main

import "fmt"

func main() {
	//Create new tasks

	tasks := []Task{
		&EmailTask{Email: "barb@mail.com", Subject: "test1", MessageBody: "testing workerpool implementation"},
		&ImgProcessingTask{ImgUrl: "/images/sample1.jpg"},
		&EmailTask{Email: "barb@mail.com", Subject: "test2", MessageBody: "testing workerpool implementation"},
		&ImgProcessingTask{ImgUrl: "/images/sample2.jpg"},
		&EmailTask{Email: "barb@mail.com", Subject: "test3", MessageBody: "testing workerpool implementation"},
		&ImgProcessingTask{ImgUrl: "/images/sample3.jpg"},
		&EmailTask{Email: "barb@mail.com", Subject: "test4", MessageBody: "testing workerpool implementation"},
		&ImgProcessingTask{ImgUrl: "/images/sample4.jpg"},
	}

	// tasks := make([]Task, 20)
	// for i := 0; i < 20; i++ {
	// 	tasks[i] = Task{ID: i + 1}
	// }

	//Create a WorkerPool
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5, //no.of workers that can run at a time
	}

	//Run the pool
	wp.Run()
	fmt.Println("All processes have been processed!")
}
