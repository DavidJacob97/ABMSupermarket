package main

var allTills = make([]int, 6)

type Till struct {
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

func getTills() []int {
	return allTills
}

func main() {

}
