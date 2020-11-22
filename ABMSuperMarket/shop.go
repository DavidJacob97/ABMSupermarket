package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Shop struct {
	timeOfDay              int
	maxCapacity            int
	handSanitizerRemaining int
	daysRemaining          int
	tills                  []Till
}

var shop Shop

type Customer struct {
	name        string
	patience    int
	isAntiMask  bool
	items       int
	queueNumber int
}

func addCustomer(setItems int, setQueueNumber int) *Customer {
	temp := Customer{items: setItems}
	temp.queueNumber = setQueueNumber
	return &temp
}

func newCustomer() {
	//customerSlice := make([]Customer, 20)
}

func timeLoop() {
	shop.timeOfDay = shop.timeOfDay + 1
	if shop.timeOfDay == 1440 {
		shop.timeOfDay = 0
		shop.daysRemaining = shop.daysRemaining - 1
	}
}

func setCovid() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 5
	covidlv := rand.Intn((max - min + 1) + min)

	switch covidlv {
	case 1:
		shop.maxCapacity = 100
	case 2:
		shop.maxCapacity = 75
	case 3:
		shop.maxCapacity = 50
	case 4:
		shop.maxCapacity = 25
	case 5:
		shop.maxCapacity = 10
	default:
		shop.maxCapacity = 100
	}
}

func openShop() {
	fmt.Println("Tills opening")

	if shop.daysRemaining == 0 {
		setCovid()
		shop.daysRemaining = 7
	}

	for shop.timeOfDay < 1320 && shop.timeOfDay >= 540 {
		customer()
		handSanitizer()
		timeLoop()
		time.Sleep(time.Second)
	}

	if shop.timeOfDay == 1320 {
		closeShop()
	}
}

func customer() {
	shop.handSanitizerRemaining = shop.handSanitizerRemaining - 1
}

func closeShop() {
	fmt.Println("no more customers allowed")
	fmt.Println("processremaining customers")
	fmt.Println("close all tills")
	fmt.Println("close shop")

	for shop.timeOfDay >= 1320 || shop.timeOfDay < 540 {
		fmt.Println("shop is closed")
		timeLoop()
		time.Sleep(time.Second)
	}

	if shop.timeOfDay == 540 {
		openShop()
	}

}

func handSanitizer() {
	if shop.handSanitizerRemaining == 0 {
		fmt.Println("Refilling hand sanitizer")
		shop.handSanitizerRemaining = 100
	}

}

type Queue struct {
	itemProcessingTime int
	inQueue            []Customer
}

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

func main() {
	shop.timeOfDay = 540
	shop.handSanitizerRemaining = 100
	fmt.Println(shop.maxCapacity)
	openShop()
}
