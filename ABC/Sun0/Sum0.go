package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	pairs := make([]struct{ a, b int }, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&pairs[i].a, &pairs[i].b)
	}

	findXWithSumZero(pairs)
}

func findXWithSumZero(pairs []struct{ a, b int }) {
	var sum int
	for _, pair := range pairs {
		sum += pair.a + pair.b
	}

	if sum != 0 {
		fmt.Println("No")
		return
	}

	fmt.Println("Yes")
	var results []int
	for _, pair := range pairs {
		for x := pair.a + 1; x < pair.b; x++ {
			results = append(results, x)
		}
	}

	for i, result := range results {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(result)
	}
	fmt.Println()
}
