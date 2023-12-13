package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	// v2.0
	// log before & after
	/*
		log.Println("Before invocation")
		add(100, 200)
		log.Println("After invocation")

		log.Println("Before invocation")
		subtract(100, 200)
		log.Println("After invocation")
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	// v3.0
	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
	logOperation(func(x, y int) {
		fmt.Println("Multiply result :", x*y)
	}, 100, 200)
}

// v3.0
func logOperation(operFn func(int, int), x, y int) {
	log.Println("Before invocation")
	operFn(x, y)
	log.Println("After invocation")
}

// v2.0
func logAdd(x, y int) {
	log.Println("Before invocation")
	add(x, y)
	log.Println("After invocation")
}

func logSubtract(x, y int) {
	log.Println("Before invocation")
	subtract(x, y)
	log.Println("After invocation")
}

// v1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
