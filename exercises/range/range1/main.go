// range1
// Make me compile!
package main

import "fmt"

func main() {
	even_numbers := []int{2, 4, 6, 8, 10}

	for n, v := range even_numbers {
		fmt.Printf("%d %d is even\n", n, v)
	}
}
