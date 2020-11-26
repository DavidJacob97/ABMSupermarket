package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// UNUSED work around
func UNUSED(x ...interface{}) {}

var mutex = &sync.Mutex{}

var foreNames = []string{"Brian", "Evan", "Martin", "Robert"}
var surNames = []string{"Hogarty", "Callaghan", "Miller", "Robson"}

// Tills is global
var Tills [6]Till

//Shop works of the time, handsanitizer,
type Shop struct {
	timeOfDay              float64
	maxCapacity            int
	handSanitizerRemaining int
	daysRemaining          int
	tills                  []Till
	customerInstore        int
	shopOpened             bool
}

// ShopStat gives back info about how well our shop is doing
type ShopStat struct {
	waitTimes                  []float64
	totalProductsProcessed     int
	averagecustomerwaitTime    float64
	averagecheckoututilisation float64
	averageproductspertrolley  int
	Thenumberoflostcustomers   int
}

var shop Shop
var stat ShopStat

//Customer object, possibly not enter cause of a mask, carries items
type Customer struct {
	name     string
	patience int
	hasMask  bool
	items    int
}

var arrivingCustomers []Customer
var customersInShop []Customer

func randomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn((max - min + 1) + min)
	return randNum
}

func addCustomerToShop() {
	for {
		//this part will remove the customer if he is not wearing mask
		if (len(arrivingCustomers) > 0) {
			mutex.Lock()
			randNum := randomNumber(0, len(arrivingCustomers) - 1)
			if arrivingCustomers[randNum].hasMask == true {
				customersInShop = append(customersInShop, arrivingCustomers[randNum])

				fmt.Printf("Customer %s has entered the shop\n", arrivingCustomers[randNum].name)

				copy(arrivingCustomers[randNum:], arrivingCustomers[randNum + 1:])
				e := Customer{}
				arrivingCustomers[len(arrivingCustomers) - 1] = e
				arrivingCustomers = arrivingCustomers[:len(arrivingCustomers) - 1]

				shop.handSanitizerRemaining -= 1
			} else {
				fmt.Printf("Customer %s does not have a mask and was refused entry\n", arrivingCustomers[randNum].name)
				copy(arrivingCustomers[randNum:], arrivingCustomers[randNum + 1:])
				e := Customer{}
				arrivingCustomers[len(arrivingCustomers) - 1] = e
				arrivingCustomers = arrivingCustomers[:len(arrivingCustomers) - 1]
			}
			mutex.Unlock()
			
		}
		//every 5 sec a customer will be added to the shop
		time.Sleep(5 * time.Second)
	}
}

func randomPause(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max*1000)))
}

func timeLoop() {
	for {
		if shop.timeOfDay == 540 {
			shop.shopOpened = true
		}
		if shop.timeOfDay == 1320 {
			shop.shopOpened = false
		}

		shop.timeOfDay = shop.timeOfDay + 1
		time.Sleep(5 * time.Millisecond)

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
	if shop.shopOpened == true {
		fmt.Println("Tills opening")
		for shop.shopOpened == true {

			if shop.daysRemaining == 0 {
				setCovid()
				shop.daysRemaining = 7
			}
			customer()
			handSanitizer()

			time.Sleep(5 * time.Millisecond)

		}

	}
	fmt.Println("no more customers allowed")
	fmt.Println("processremaining customers")
	fmt.Println("close all tills")
	fmt.Println("close shop")

	shop.timeOfDay = 540

}

func customer() {
	shop.handSanitizerRemaining = shop.handSanitizerRemaining - 1
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

func getAvgQueueLength() int {

	allQueues := Tills
	totalQueueLength := 0

	for i := 0; i < 6; i++ {
		totalQueueLength = len(allQueues[i].queue.inQueue)
	}

	avgQueueLength := totalQueueLength / len(Tills)
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
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}

// customer chooses the fast track if it is open and if he has less than 15 items
func fastTrackOrStandard(c Customer) {
	if Tills[0].isFastTrack == true {
		if c.items < 15 {
			Tills[0].queue.inQueue = append(Tills[0].queue.inQueue, c)

		}
	}
}

// locates the shortest queue out of the 5 standard tills using Tills var
func findBestTill() Till {
	shortestQueue := 1
	for i := 1; i < len(Tills)-1; i++ {
		if Tills[i].isOpen {
			for j := 1; j < len(Tills)-1; j++ {
				if Tills[j].isOpen {
					if len(Tills[i].queue.inQueue) < len(Tills[j].queue.inQueue) {
						shortestQueue = i
					}
				}
			}
		}
	}
	return Tills[shortestQueue]
}

func generateCustomers() {
	for {
		//if shop.customerInstore < shop.maxCapacity && shop.shopOpened == true {

		r := rand.Intn(len(foreNames))
		foreName := foreNames[r]

		r = rand.Intn(len(surNames))

		lastName := surNames[r]
		name := foreName + " " + lastName

		customer := Customer{name: name}
		customer.hasMask = true //need some code in some chance
		customer.items = 5      //to be generated randomly
		customer.patience = 0   //to be generated randomly

		mutex.Lock()
		arrivingCustomers = append(arrivingCustomers, customer)
		mutex.Unlock()
		shop.customerInstore = shop.customerInstore + 1
		//}
		//generate new customer every 5 sec
		time.Sleep(time.Duration(5 * time.Second))

	}
}

func testPrintAllCustomers() {
	for {
		fmt.Println("Customers in allCustomers:")
		mutex.Lock()
		for i := 0; i < len(arrivingCustomers); i++ {
			fmt.Print(arrivingCustomers[i].name + ", ")
		}
		mutex.Unlock()
		fmt.Println()

		//print array every 10 secs
		time.Sleep(time.Duration(10 * time.Second))
	}
}

func testPrintAllCustomersInShop() {
	for {
		fmt.Println("Customers in customersInShop:")
		mutex.Lock()
		for i := 0; i < len(customersInShop); i++ {
			fmt.Print(customersInShop[i].name + ", ")
		}
		mutex.Unlock()
		fmt.Println()

		//print array every 10 secs
		time.Sleep(time.Duration(10 * time.Second))
	}
}

func processCustomer(till Till) {
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
	setCovid()
	rand.Seed(time.Now().UnixNano())
	//var wg sync.WaitGroup
	//wg.Add(2)
	go testPrintAllCustomers()
	go testPrintAllCustomersInShop()
	go generateCustomers()
	go addCustomerToShop()
	//shop.timeOfDay = 540
	//shop.handSanitizerRemaining = 100
	//go timeLoop()

	Tills[0] = *newTill("Fast track", true, true, 2)
	Tills[1] = *newTill("Till 1", false, false, 3)
	Tills[2] = *newTill("Till 2", false, false, 3)
	Tills[3] = *newTill("Till 3", false, false, 3)
	Tills[4] = *newTill("Till 4", false, false, 3)
	Tills[5] = *newTill("Till 5", false, false, 3)

	for {

		//calling processCustomer for each till for processing the customers in queue
		/*
			for i := 0; i < len(Tills); i++ {
				if Tills[i].isOpen {
					processCustomer(Tills[i])

				}

			}
		*/

		//openShop()

	}
}
