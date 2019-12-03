package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	check(err)

	moduleMasses := strings.Split(string(input), "\n")
	var fuel []int64

	for _, mass := range moduleMasses {
		if mass != "" {
			massInt, err := strconv.ParseInt(mass, 10, 64)
			check(err)
			fuel = append(fuel, calculateFuel(massInt))
		}
	}

	var total int64

	for _, amt := range fuel {
		total = total + amt
	}

	fmt.Println(total)
}

func calculateFuel(mass int64) int64 {
	return (mass / 3) - 2
}
