package truck

import (
	"fmt"
)

type Truck struct {
	Brand string
	Model string
}

func GetTruck(truck Truck) string {
	if truck == (Truck{}) {
		return "Truck is empty"
	}

	return fmt.Sprintf("Brand: %s\nModel: %s", truck.Brand, truck.Model)
}
