package main

// Customer is a person buying food
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
