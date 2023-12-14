package calculator

import "fmt"

var opCount map[string]int

func init() {
	opCount = make(map[string]int)
	fmt.Println("calculator[init.go] - initialized")
}

func OpCount() map[string]int {
	return opCount
}

func incrementOperation(operation string) {
	opCount[operation]++
}
