package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key", key, "value:", value)
	}
}
