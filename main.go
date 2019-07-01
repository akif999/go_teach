package main

import "fmt"

func main() {
	// declaration of variables
	var num int
	var name string
	var nums []int
	var names []string

	// assignment of variables
	num = 10
	name = "Jhon"
	nums = []int{1, 2, 3, 4, 5}
	names = []string{"Richard", "Eric", "Bighead", "Gavin"}

	// output variable with formats
	fmt.Printf("the number is : %d\n", num)
	fmt.Printf("the name is : %s\n\n", name)

	for _, n := range nums {
		fmt.Printf("the nunber is : %d\n", n)
	}

	for _, n := range names {
		fmt.Printf("the name is : %s\n", n)
	}
}
