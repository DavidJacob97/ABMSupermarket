package main

import (
	"fmt"
	"math/rand"
	"time"
)

var timeofday int
var maxcapacity int
var handsanitizerremaining int
var DaysRemaining int
func timeloop() {

	timeofday = timeofday + 1

	if timeofday == 1440 {

		timeofday=0
		DaysRemaining=DaysRemaining-1

	
    }
}
func setCovid() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 5
	covidlv := rand.Intn((max - min + 1) + min)


	if covidlv == 0 {

		maxcapacity = 100

	}
	if covidlv == 1 {
		maxcapacity = 100
	}

	if covidlv == 2 {

		maxcapacity = 75

	}
	if covidlv == 3 {

		maxcapacity = 50

	}

	if covidlv == 4 {

		maxcapacity = 25

	}

	if covidlv == 5 {

		maxcapacity = 10

	}

	

}

func openshop() {


     
	fmt.Println("Tills opening")
	
	if DaysRemaining==0{
	    setCovid()
	    
	    DaysRemaining=7
	}
	
	for timeofday < 1320 && timeofday >= 540 {

		customer()
		handsanitizer()
		timeloop()

	}
   if timeofday == 1320 {

		closeshop()
	}
}

func customer() {

	
	handsanitizerremaining = handsanitizerremaining - 1
}

func closeshop() {

	fmt.Println("no more customers allowed")
	fmt.Println("processremaining customers")
	fmt.Println("close all tills")
	fmt.Println("close shop")
	
		for timeofday >= 1320 || timeofday < 540 {

		fmt.Println("shop is closed")
		
      timeloop()
	
	}
	
	 if timeofday == 540 {

<<<<<<< HEAD
=======
	openshop()
	}
	
}

//func closeAllTills() {
//	for i := 0; i < len(allTills); i++ {
//		allTills[i].isOpen = false
//	}
//}

//func getTills() []Till {
//	return allTills
//}

>>>>>>> d281fe5eaa641165b5c1b3bb13dcf7d504920c72
func handsanitizer() {

	if handsanitizerremaining == 0 {

		fmt.Println("refiling hand sanitizer")
		handsanitizerremaining = 100
	}

}

func main() {
	timeofday = 540

	handsanitizerremaining = 100


	fmt.Println(maxcapacity)
    openshop()

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


