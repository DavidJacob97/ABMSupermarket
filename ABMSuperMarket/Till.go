package main

var allTills = make([]Till, 6)

//Till logic if > 6 , then add till part
type Till struct {
	till        string
	isFastTrack bool
	queueLength int
	open        bool
	closed      bool
}

func openTill() {
	allTills = allTills[:1]
}
func closeTill() {
	allTills = allTills[1:]
}

func closeAllTills() {
	allTills = allTills[:0]
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
