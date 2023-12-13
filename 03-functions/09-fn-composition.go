package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
		logOperation(func(x, y int) {
			fmt.Println("Multiply result :", x*y)
		}, 100, 200)
	*/

	// v4.0
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	// v5.0
	/*
		profiledAdd := getProfiledOperation(add)
		profiledAdd(100, 200)

		profiledSubtract := getProfiledOperation(subtract)
		profiledSubtract(100, 200)
	*/

	logAdd := getLogOperation(add)
	profiledLogAdd := getProfiledOperation(logAdd)
	profiledLogAdd(100, 200)

	profiledLogSubtract := getProfiledOperation(getLogOperation(subtract))
	profiledLogSubtract(100, 200)

	getProfiledOperation(getLogOperation(func(x, y int) {
		fmt.Println("Multiply Result :", x*y)
	}))(100, 200)
}

// v5.0
func getProfiledOperation(operFn func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operFn(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

// v4.0
func getLogOperation(operFn func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Before invocation")
		operFn(x, y)
		log.Println("After invocation")
	}
}

// v3.0
/*
func logOperation(operFn func(int, int), x, y int) {
	log.Println("Before invocation")
	operFn(x, y)
	log.Println("After invocation")
}
*/

// v1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
