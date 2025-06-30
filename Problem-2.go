package main

import "fmt"

// Problem-2: With a single integer as the input, generate the following until a = x [series of numbers as shown in below examples]

//   Output: (examples)
//     1) input a = 1, then output : 1
//     2) input a = 2, then output : 1, 3
//     3) input a = 3, then output : 1, 3, 5
//     4) input a = 4, then output : 1, 3, 5, 7
//     .
//     .
//     5) input a = x, then output : 1, 3, 5, 7, .......

func main() {
	x := 0
	fmt.Print("Enter the value x :")
	fmt.Scan(&x)
	result := Problem2(x)
	fmt.Println("Result : ", result)
}

func Problem2(input int) []int {
	odd := 1
	result := []int{}
	for i := 0; i < input; i++ {
		// fmt.Println("i :", odd)
		result = append(result, odd)
		odd += 2
	}
	return result
}
