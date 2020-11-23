package main

import (
	"fmt"
	"math/rand"
	"time"
)

<<<<<<< HEAD
var foreNames = []string{"Brian", "Evan", "Martin", "Robert"}
var surNames = []string{"Hogarty", "Callaghan", "Miller", "Robson"}
=======
func init() {
	// Random seed
	rand.Seed(time.Now().UTC().UnixNano())
}
>>>>>>> de003e7eb72b162ab7dd719d10f078bf5e9af0c1

//Shop works of the time, handsanitizer,
type Shop struct {
	timeOfDay              int
	maxCapacity            int
	handSanitizerRemaining int
	daysRemaining          int
	tills                  []Till
}

var shop Shop

//Customer has patience var, possibly not enter cause of a mask, carries items
type Customer struct {
	name       string
	patience   int
	isAntiMask bool
	items      int
}

// hopefully finaly everithing work
func (c *Customer) addCustomer(setItems int) {
	c.items = setItems
}

func newCustomer() {
	//customerSlice := make([]Customer, 20)
}

func randomPause(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max*1000)))
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

	fmt.Println(shop.maxCapacity)
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
		time.Sleep(5 * time.Millisecond)
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
		time.Sleep(5 * time.Millisecond)
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

// Queue has processing item time and Customers array queue
type Queue struct {
	itemProcessingTime int
	inQueue            []Customer
}

//Till has a type, can be opened and contains a queue
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

func getAvgQueueLength(q1 Queue, q2 Queue, q3 Queue, q4 Queue, q5 Queue, q6 Queue) int {

	allQueues := []Queue{q1, q2, q3, q4, q5, q6}
	totalQueueLength := 0
	for i := 0; i < 6; i++ {
		totalQueueLength = len(allQueues[i].inQueue)
	}
	avgQueueLength := totalQueueLength / len(allQueues)

	return avgQueueLength

}

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

func generateCustomers() {
	//for {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(foreNames))
	foreName := foreNames[r]

	r = rand.Intn(len(surNames))

	lastName := surNames[r]
	name := foreName + " " + lastName

	fmt.Println(name)

	//}
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
	generateCustomers()

	shop.timeOfDay = 540
	shop.handSanitizerRemaining = 100
	//openShop()
}
