package main

import (
	"fmt"

	"github.com/carloshdurante/bdd-golang/truck"
)

func main() {
	truck := truck.Truck{
		Brand: "Ford",
		Model: "F150",
	}

	fmt.Println(truck.GetTruck(truck))
}
