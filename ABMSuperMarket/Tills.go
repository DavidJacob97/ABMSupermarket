package main

var allTills = make([]int, 5)

type tills struct {
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

func main() {

}
