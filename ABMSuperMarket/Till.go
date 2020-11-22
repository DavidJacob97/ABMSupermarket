package main

//5 standard, 1 fast till
var allTills = make([]Till, 6)

//Till logic if > 6 , then add till part
type Till struct {
	till        string
	isFastTrack bool
	queueLength int
	open        bool
	closed      bool
}

func openTill(t Till) {
	t.open = true
	t.closed = false
}
func closeTill(t Till) {
	t.open = false
	t.closed = true
}

func closeAllTills() {
	for i := 0; i < len(allTills); i++ {
		allTills[i].closed = true
		allTills[i].open = false
	}
}

func getTills() []Till {
	return allTills
}

func getAvgQueueNum() int {
	sum := 0
	for i := 0; i < len(allTills); i++ {
		sum += allTills[i].queueLength
	}
	avgQueueLength := sum / len(allTills)
	return avgQueueLength
}
