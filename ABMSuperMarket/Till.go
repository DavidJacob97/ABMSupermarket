package main

import (
	"fmt"
	"time"
)

//5 standard, 1 fast till
var allTills = make([]Till, 6)

//Queue has a till and item processing time and an array
type Queue struct {
	itemProcessingTime int
	inQueue            []Customer
}

//Till logic if > 6 , then add till part
type Till struct {
	name        string
	isFastTrack bool
	isOpen      bool
	queue       Queue
}

func openTill(t Till) {
	t.isOpen = true
}
func closeTill(t Till) {
	t.isOpen = false
}

func closeAllTills() {
	for i := 0; i < len(allTills); i++ {
		allTills[i].isOpen = false
	}
}

func getTills() []Till {
	return allTills
}

/*
func getAvgQueueNum() int {
	sum := 0
	for i := 0; i < len(allTills); i++ {
		sum += allTills[i].queueLength
	}
	avgQueueLength := sum / len(allTills)
	return avgQueueLength
}
*/

func newQueue(itemProcessingTime int) *Queue {
	q := Queue{}
	q.itemProcessingTime = itemProcessingTime
	return &q
}

func newTill(name string, isFastTrack bool, isOpen bool, itemProcessingTime int) *Till {
	t := Till{name: name}
	t.isFastTrack = isFastTrack
	t.isOpen = isOpen
	t.queue = *newQueue(itemProcessingTime)
	return &t
}

func remove(slice []Customer, i int) []Customer {
	return append(slice[:i], slice[i+1:]...)
}

func processCustomer(till Till) {
	if !till.isOpen {
		fmt.Printf("Till %s is currently closed\n", till.name)
		return
	}

	if len(till.queue.inQueue) == 0 {
		fmt.Printf("No customers currently in queue at till %s\n", till.name)
		return
	}

	//process the first customer in queue
	fmt.Printf("Processing customer %s at till %s\nNumber of items: %d\n", till.queue.inQueue[0].name, till.name, till.queue.inQueue[0].items)

	for i := till.queue.inQueue[0].items; i != 0; i-- {
		time.Sleep(time.Duration(till.queue.itemProcessingTime) * time.Second)
		fmt.Printf("Processed item %d\n", i)
	}

	fmt.Printf("Customer %s has checked out\n", till.queue.inQueue[0].name)

	//remove first element in customer slice from queue and maintain order
	till.queue.inQueue = remove(till.queue.inQueue, 0)
}
