// maps2
// Make me compile!
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["John"] = 30
	m["Ana"] = 30
	fmt.Printf("John is %d and Ana is %d", m["John"], m["Ana"])
}
