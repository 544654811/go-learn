package demo

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id     int
	Number int
}

type Result struct {
	job *Job
	sum int
}

func cals(job *Job, resultChan chan *Result) {
	sum := 0
	number := job.Number
	for number > 0 {
		temp := number % 10
		sum += temp
		number = number / 10
	}
	resulet := &Result{
		job: job,
		sum: sum,
	}
	resultChan <- resulet
}

func Worker(jobChan chan *Job, resultChan chan *Result) {
	for job := range jobChan {
		cals(job, resultChan)
	}
}

func StartWorker(number int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < number; i++ {
		go Worker(jobChan, resultChan)
	}
}

func PrintResult(resultChan chan *Result) {
	for result := range resultChan {
		fmt.Printf("id=%v, number=%v, result=%d \n", result.job.Id, result.job.Number, result.sum)
	}
}

func Test() {
	jobChan := make(chan *Job, 1000)
	resultChan := make(chan *Result, 1000)
	StartWorker(128, jobChan, resultChan)
	go PrintResult(resultChan)

	id := 1
	for {
		id++
		number := rand.Int()

		jobChan <- &Job{
			Id:     id,
			Number: number,
		}
	}
}
