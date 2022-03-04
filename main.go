package main

import (
	"fmt"

	t "github.com/carloshdurante/bdd-golang/truck"
)

func main() {
	truck := t.Truck{
		Brand: "Ford",
		Model: "F150",
	}

	fmt.Println(t.GetTruck(truck))
}
