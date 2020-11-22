package main

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

func processCustomer(Customer) {

}
