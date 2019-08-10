package dynamicProgramming

import (
	"fmt"
)

// https://leetcode-cn.com/problems/climbing-stairs/

func ClimbStairs()  {
	fmt.Println("Climbing Stairs: ")

	for _, targetN := range []int{10,20,30,40,50,100} {
		res := climbStairs(targetN)
		fmt.Printf("Climbing %d stairs has %d ways\n", targetN, res)
	}

	fmt.Println()
}

func climbStairs(n int) int {
	if n < 1 {
		panic(fmt.Sprintf("`n` must be a positive int, but got: %d", n))
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	a := 1
	b := 2
	var temp int

	for i := 2; i < n; i++ {
		temp = a + b
		a = b
		b = temp
	}

	return temp
}

