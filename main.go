package main

import "fmt"

type character struct {
	name        string
	age         int
	devilFruit  string
	affiliation string
}

func main() {
	luffy := character{name: "Monkey D. Luffy", devilFruit: "Gomu Gomu no Mi", affiliation: "Straw Hat Pirates"}
	zoro := character{name: "Roronoa Zoro", age: 17, affiliation: "Straw Hat Pirates"}

	fmt.Printf("Name: %s Devil fruit: %s Affiliation: %s \n", luffy.name, luffy.devilFruit, luffy.affiliation)
	fmt.Printf("Name: %s Age: %d Affiliation: %s", zoro.name, zoro.age, zoro.affiliation)

}
