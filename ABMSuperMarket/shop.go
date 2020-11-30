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
var daysOfSimulation int = 1
var foreNames = []string{"Brian", "Evan", "Martin", "Robert"}
var surNames = []string{"Hogarty", "Callaghan", "Miller", "Robson"}

// Tills is global variable
var Tills [6]Till

type Shop struct {
	timeOfDay              float64
	maxCapacity            int
	handSanitizerRemaining int
	daysRemaining          int
	tills                  []Till
	customerInStore        int
	shopOpened             bool
	scanTime               float64
}

// ShopStat gives back info about how well our shop is doing
type ShopStat struct {
	waitTimes                  []int64
	totalProductsProcessed     int
	averageCustomerWaitTime    int64
	averageCheckoutUtilisation float64
	averageProductsPerTrolley  int
	theNumberOfLostCustomers   int
}

var shop Shop
var stat ShopStat

//Customer object, possibly not enter cause of a mask, carries items
type Customer struct {
	name    string
	hasMask bool
	items   int
	arrival int64
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
		if len(arrivingCustomers) > 0 && len(customersInShop) < shop.maxCapacity && shop.shopOpened == true {
			mutex.Lock()
			randNum := randomNumber(0, len(arrivingCustomers)-1)
			if arrivingCustomers[randNum].hasMask == true {
				arrivingCustomers[randNum].arrival = makeTimestamp()
				customersInShop = append(customersInShop, arrivingCustomers[randNum])
				fmt.Printf("Customer %s has entered the shop\n", arrivingCustomers[randNum].name)

				copy(arrivingCustomers[randNum:], arrivingCustomers[randNum+1:])
				e := Customer{}
				arrivingCustomers[len(arrivingCustomers)-1] = e
				arrivingCustomers = arrivingCustomers[:len(arrivingCustomers)-1]

				shop.handSanitizerRemaining--
			} else {
				fmt.Printf("Customer %s does not have a mask and was refused entry\n", arrivingCustomers[randNum].name)

				copy(arrivingCustomers[randNum:], arrivingCustomers[randNum+1:])
				e := Customer{}
				arrivingCustomers[len(arrivingCustomers)-1] = e
				arrivingCustomers = arrivingCustomers[:len(arrivingCustomers)-1]
			}
			mutex.Unlock()

		}
		//every 10 sec a customer will be added to the shop
		time.Sleep(300 * time.Millisecond)
	}
}

func randomPause(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max*1000)))
}

func timeLoop() {
	for {
		shop.shopOpened = true
		println("Shop Is open,Welcome ")
		time.Sleep(60 * time.Second)
		shop.shopOpened = false
		println("shop is closed no more people allowed to enter")
		time.Sleep(60 * time.Second)
		daysOfSimulation--
	}
}

func setCovid() {
	level := randomNumber(1, 5)

	switch level {
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
	fmt.Printf("Current Covid-19 Restriction Level: %d\nMax Customers Allowed to Enter: %d", level, shop.maxCapacity)
}

func handSanitizer() {
	if shop.handSanitizerRemaining == 0 {
		fmt.Println("Refilling hand sanitizer")
		shop.handSanitizerRemaining = 100
	}
}

// Queue has processing item time and Customers array queue
type Queue struct {
	isBusy             bool
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
	totalQueueLength := 0
	totalOpenTills := 0
	for i := 1; i < len(Tills); i++ {
		if Tills[i].isOpen {
			totalOpenTills += 1
			totalQueueLength += len(Tills[i].queue.inQueue)
		}
	}
	avgQueueLength := totalQueueLength / totalOpenTills
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

func fastTrackOrStandard(c Customer) {
	if Tills[0].isFastTrack == true {
		if c.items < 15 {
			Tills[0].queue.inQueue = append(Tills[0].queue.inQueue, c)
		}
	}
}

// Locates the shortest queue out of the 5 standard tills using Tills var
func findBestTill() *Till {
	shortestQueue := 1
	for i := 1; i < len(Tills); i++ {
		if Tills[i].isOpen {
			for j := 1; j < len(Tills); j++ {
				if Tills[j].isOpen {
					if len(Tills[i].queue.inQueue) < len(Tills[j].queue.inQueue) {
						shortestQueue = i
					}
				}
			}
		}
	}
	return &Tills[shortestQueue]
}

func makeTimestamp() int64 {
	now := time.Now()
	return now.Unix()
}

func generateCustomers() {
	for daysOfSimulation > 0 {
		if len(arrivingCustomers) < (shop.maxCapacity * 2) {
			r := rand.Intn(len(foreNames))
			foreName := foreNames[r]

			name := foreName

			customer := Customer{name: name}
			if randomNumber(1, 60)%3 == 2 {
				customer.hasMask = false
			} else {
				customer.hasMask = true
			}
			customer.items = randomNumber(1, 200) //to be generated randomly

			mutex.Lock()
			arrivingCustomers = append(arrivingCustomers, customer)
			mutex.Unlock()
			shop.customerInStore += 1

			time.Sleep(time.Duration(200 * time.Millisecond))
			//if shop.customerInStore > 10 {
			//	fmt.Printf("Exit\n")
			//	return
			//}
		}
		//generate new customer every 5 sec
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
		time.Sleep(10 * time.Second)
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
		time.Sleep(10 * time.Second)
	}
}

func processItems(i int) {
	Tills[i].queue.isBusy = true

	for j := Tills[i].queue.inQueue[0].items; j != 0; j-- {
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Printf("Customer %s has checked out\n", Tills[i].queue.inQueue[0].name)

	//Remove first element in customer slice from queue and maintain order
	mutex.Lock()

	stat.totalProductsProcessed = stat.totalProductsProcessed + Tills[i].queue.inQueue[0].items
	stat.waitTimes = append(stat.waitTimes, (makeTimestamp() - Tills[i].queue.inQueue[0].arrival))
	//fmt.Printf("Secs: %d %d %d\n", makeTimestamp(), Tills[i].queue.inQueue[0].arrival, (makeTimestamp() - Tills[i].queue.inQueue[0].arrival))

	copy(Tills[i].queue.inQueue[0:], Tills[i].queue.inQueue[1:])
	e := Customer{}
	Tills[i].queue.inQueue[len(Tills[i].queue.inQueue)-1] = e
	Tills[i].queue.inQueue = Tills[i].queue.inQueue[:len(Tills[i].queue.inQueue)-1]
	Tills[i].queue.isBusy = false
	mutex.Unlock()
}

func getAvgItems(x []int) float64 {
	n := len(x)
	sum := 0
	for i := 0; i < n; i++ {
		sum += (x[i])
	}
	avg := (float64(sum)) / (float64(n))
	return avg
}

func getAvgTimes(x []int64) int64 {
	n := len(x)
	var sum int64 = 0
	for i := 0; i < n; i++ {
		sum += x[i]
	}
	avg := sum / int64(n)
	return avg
}

func processCustomers() {
	for {

		for i := 0; i < len(Tills); i++ {
			if Tills[i].isOpen {
				if len(Tills[i].queue.inQueue) == 0 {
					fmt.Printf("No customers currently in queue at %s\n", Tills[i].name)
					continue
				}

				//process the first customer in queue
				fmt.Printf("Processing customer %s at %s\n\n",
					Tills[i].queue.inQueue[0].name, Tills[i].name)

				if !Tills[i].queue.isBusy {
					go processItems(i)
				}

				avgQueueLength := getAvgQueueLength()
				if avgQueueLength < 2 {
					totalOpenTills := 0
					for i := 1; i < len(Tills); i++ {
						if Tills[i].isOpen {
							totalOpenTills += 1
						}
					}
					fmt.Printf("totalOpenTills %d\n", totalOpenTills)
					if totalOpenTills != 1 {
						Tills[i].isOpen = false
						fmt.Printf("Closing %s\n", Tills[i].name)
					}
				}
			} else {
				if len(Tills[i].queue.inQueue) != 0 {
					if !Tills[i].queue.isBusy {
						go processItems(i)
					}
				}
				//We open another till if average queue length for each open till is greater than 5
				avgQueueLength := getAvgQueueLength()
				if avgQueueLength > 5 {
					fmt.Printf("Opening %s\n", Tills[i].name)
					Tills[i].isOpen = true
				}
			}
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func printStat() {
	stat.averageCustomerWaitTime = getAvgTimes(stat.waitTimes)
	stat.averageProductsPerTrolley = int(stat.totalProductsProcessed / len(stat.waitTimes))
	println(stat.totalProductsProcessed)
	fmt.Printf("Average Customer Wait Time: %d\nSize: %d\n", stat.averageCustomerWaitTime, len(stat.waitTimes))
	println(stat.averageCheckoutUtilisation)
	println(stat.averageProductsPerTrolley)
	println(stat.theNumberOfLostCustomers)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	setCovid()
	go testPrintAllCustomers()
	go testPrintAllCustomersInShop()
	go generateCustomers()
	go addCustomerToShop()
	go processCustomers()
	//shop.timeOfDay = 540
	shop.handSanitizerRemaining = 50
	go timeLoop()

	Tills[0] = *newTill("Fast track", true, true, 1)
	Tills[1] = *newTill("Till 1", false, true, 1)
	Tills[2] = *newTill("Till 2", false, false, 1)
	Tills[3] = *newTill("Till 3", false, false, 1)
	Tills[4] = *newTill("Till 4", false, false, 1)
	Tills[5] = *newTill("Till 5", false, false, 1)

	for daysOfSimulation > 0 {
		if len(customersInShop) > 0 {
			randNum := randomNumber(0, len(customersInShop)-1)
			mutex.Lock()
			c := customersInShop[randNum]
			if c.items > 6 {
				t := findBestTill()

				t.queue.inQueue = append(t.queue.inQueue, c)
				fmt.Printf("Customer %s has entered the queue at %s\n", c.name, t.name)

				copy(customersInShop[randNum:], customersInShop[randNum+1:])
				e := Customer{}
				customersInShop[len(customersInShop)-1] = e
				customersInShop = customersInShop[:len(customersInShop)-1]
			} else {
				fmt.Printf("Customer %s has entered the fast track queue\n", c.name)

				Tills[0].queue.inQueue = append(Tills[0].queue.inQueue, c)

				copy(customersInShop[randNum:], customersInShop[randNum+1:])
				e := Customer{}
				customersInShop[len(customersInShop)-1] = e
				customersInShop = customersInShop[:len(customersInShop)-1]
			}
			mutex.Unlock()
		} else {
			//fmt.Println("Waiting for customers\n")
		}
		time.Sleep(500 * time.Millisecond)

	}
	printStat()
}
