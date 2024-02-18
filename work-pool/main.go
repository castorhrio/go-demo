package main

import (
	"fmt"
	"strconv"
)

func main() {
	tasks := make([]Task, 0, 30)
	for i := 0; i < 30; i++ {
		if i%2 == 0 {
			tasks = append(tasks, &EmailTask{
				Email:   strconv.Itoa(i+1) + "@gmail.com",
				Subject: strconv.Itoa(i + 1),
				Message: "this is message: " + strconv.Itoa(i+1),
			})
		} else {
			tasks = append(tasks, &ImageTask{
				ImageUrl: "this is image url " + strconv.Itoa(i+1),
			})
		}
	}

	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5,
	}

	wp.Run()
	fmt.Printf("All tasks have been processed")
}
