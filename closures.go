/*
Go functions may be closures. A closure is a function value that references variables from outside its body
*/
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func muller() func(i int) int {
	x := 2
	return func(a int) int {
		x *= a
		return x
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	mul := muller()
	fmt.Println(mul(20))
}
