package main

import "fmt"

type player struct {
	name string
	age  int
}

func main() {
	player := player{name: "Victor", age: 19}
	fmt.Println(player)
	fmt.Println(player.name)
	fmt.Println(player.age)
}
