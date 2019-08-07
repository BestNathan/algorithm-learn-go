package concurrency

import (
	"fmt"
)

func Run()  {
	fmt.Println()
	fmt.Println("================= concurrency ==================")

	fmt.Println("random order print: ")
	RandOrderPrint()
	fmt.Println()
	
	fmt.Println("order print: ")
	OrderPrint()

	fmt.Println("================= concurrency ==================")
}