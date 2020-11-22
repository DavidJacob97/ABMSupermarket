package main

import (
	"fmt"
	"time"
)

type Queue struct {
	itemProcessingTime int
	till               Till
	inQueue            []Customer
}

func newQueue(till Till, itemProcessingTime int) *Queue {
	q := Queue{}
	q.itemProcessingTime = 3
	q.till = till
	return &q
}

func remove(slice []Customer, i int) []Customer {
	return append(slice[:i], slice[i+1:]...)
}

func processCustomer(queue Queue) {
	if !queue.till.open {
		fmt.Printf("Till %s is currently closed\n", queue.till.till)
		return
	}

	if len(queue.inQueue) == 0 {
		fmt.Printf("No customers currently in queue at till %s\n", queue.till.till)
		return
	}

	//process the first customer in queue
	fmt.Printf("Processing customer %s at till %s\nNumber of items: %d\n", queue.inQueue[0].name, queue.till.till, queue.inQueue[0].items)

	for i := queue.inQueue[0].items; i != 0; i-- {
		time.Sleep(time.Duration(queue.itemProcessingTime) * time.Second)
		fmt.Printf("Processed item %d\n", i)
	}

	fmt.Printf("Customer %s has checked out\n", queue.inQueue[0].name)

	//remove first element in customer slice from queue and maintain order
	queue.inQueue = remove(queue.inQueue, 0)
}
