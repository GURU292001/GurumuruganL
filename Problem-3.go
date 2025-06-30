package main

import "fmt"

// Problem-3: With a single integer as the input, generate the following until a = x [series of numbers as shown in below examples]

//   Output: (examples)
//     1) input a = 1, then output : 1
//     2) input a = 2, then output : 1
//     3) input a = 3, then output : 1, 3, 5
//     4) input a = 4, then output : 1, 3, 5
//     5) input a = 5, then output : 1, 3, 5, 7, 9
//     6) input a = 6, then output : 1, 3, 5, 7, 9
//     .
//     .
//     7) input a = x, then output : 1, 3, 5, 7, .......

func main() {
	x := 0
	fmt.Print("Enter the value x :")
	fmt.Scan(&x)
	if x%2 == 0 {
		x--
	}
	Result := Problem3(x)
	fmt.Println("Result :", Result)
}

func Problem3(input int) []int {
	result := []int{}
	odd := 1
	for i := 0; i < input; i++ {
		result = append(result, odd)
		odd += 2

	}
	return result
}
