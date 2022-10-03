package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)
func main() {
	var peoplecount int
	var name string
	var people [] string
	fmt.Println("How many people? ")
	fmt.Scanln(&peoplecount)
	groupsize := peoplecount / 2
	if peoplecount%2 != 0 {
		fmt.Println("You need an even number of people to make groups")
		os.Exit(0)
	}
	for i := 1; i <= peoplecount; i++ {
		fmt.Println("Enter name ")
		fmt.Scanln(&name)
		people = append(people, name)
    }
	for i := 1; i <= groupsize; i++ {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(people), func(i, j int) { people[i], people[j] = people[j], people[i] })
		for e := 0; e < len(people); e++ {
		fmt.Println("Group ", i, people[0], " buys for ", people[1])
		people = people[2:]
  }
		}
	}