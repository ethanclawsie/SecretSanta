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
	var budget int
	var people [] string
	var arrone [] string
	var arrtwo [] string

	x := 0
	y := 1

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

	fmt.Println("Enter budget ")
	fmt.Scanln(&budget)

	for i := 1; i <= groupsize; i++ {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(people), func(i, j int) { people[i], people[j] = people[j], people[i] })
	}
	
	fmt.Println("***************")
	fmt.Println("Budget: ",budget)
	fmt.Println("***************")
	for e := 0; e < groupsize; e++ {
		fmt.Println(people[x], " buys for ", people[x+1])
		arrone = append(arrone, people[y])
		arrtwo = append(arrtwo, people[x])
		x += 2
		y += 2
	}
  for a := 0; a < groupsize; a++ {
		randnumone := rand.Intn(len(arrone))
		randnumtwo := rand.Intn(len(arrtwo))
		fmt.Println(arrone[randnumtwo], " buys for ",arrtwo[randnumone])
		arrtwo = append(arrtwo[:randnumone], arrtwo[randnumone+1:]...)
		arrone = append(arrone[:randnumtwo], arrone[randnumtwo+1:]...)
    }
		fmt.Println("***************")
}