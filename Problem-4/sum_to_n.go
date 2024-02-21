package main

import "fmt"

func sum_to_n_a(n int) int {
		// your code here
		// O(n) recursion 
		// time complexity: O(n) 
		// it takes n recursive calls to add all numbers together
		// each recursive call takes O(1) time
		// space complexity: O(n) to store the n recursive calls on stack
		if n == 1 {
			return 1
		} else {
			return sum_to_n_a(n - 1) + n
		}
	}

	func sum_to_n_b(n int) int {
		// your code here
		// for loop to add all n numbers together
		// time complexity: O(n) 
		// It takes n iterations to add all numbers together, each iteration takes O(1) time
		// space complexity: O(1) 
		// (only the input n and extra variables like sum and i are used)
		sum := 0
		for i := 1; i <= n; i++ {
			sum += i
		}
		return sum
	}

	func sum_to_n_c(n int) int {
		// your code here
		// Use the summation formula
		// time complexity: O(1) 
		// It takes O(1) time to calculate the sum using the formula
		// space complexity: O(1)
		// no extra space other than variable n and output sum
		return n * (n + 1) / 2
	}

func main() {
	for i := 1; i <= 100; i+=5 {
		fmt.Println(sum_to_n_a(i) == sum_to_n_b(i) && sum_to_n_b(i) == sum_to_n_c(i))
	}
}