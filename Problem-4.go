package main

import "fmt"

// Problem-4: Get the total count of number listed in the dictionary which is multiple of [1,2,3,4,5,6,7,8,9]
//   (examples)
//   input: [1,2,8,9,12,46,76,82,15,20,30]
//   Output:
//     {1: 11, 2: 8, 3: 4, 4: 4, 5: 3, 6: 2, 7: 0, 8: 1, 9: 1}

func main() {
	Input := []int{1, 2, 8, 9, 12, 46, 76, 82, 15, 20, 30}
	Result := Problem4(Input)
	fmt.Println("Result :", Result)
}

func Problem4(array []int) map[int]int {
	result := map[int]int{}
	for i := 1; i < 10; i++ {
		count := 0
		for j := 0; j < len(array); j++ {
			if array[j]%i == 0 {
				count++
				result[i] = count
			}
		}
		if count == 0 {
			result[i] = 0
		}
	}
	return result

}
