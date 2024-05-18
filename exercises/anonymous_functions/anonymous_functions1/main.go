// anonymous functions1

package main

import "fmt"

func main() {

	func(name string) {
		fmt.Printf("Hello %s", name)
	}("adam")

}
