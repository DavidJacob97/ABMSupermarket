package main

import "fmt"

type Queue struct {
	count              int
	itemProcessingTime int
	till               Till
	inQueue            []Customer
}

func newQueue(till Till, itemProcessingTime int) *Queue {
	q := Queue{}
	q.count = 0
	q.itemProcessingTime = 3
	q.till = till
	return &q
}

func processCustomer(queue Queue) {
	//process the first customer in queue
	fmt.Printf("Processing customer %s at till %s\nNumber of items: %d\n", queue.inQueue[0].name, queue.till.till, queue.inQueue[0].items)

}
