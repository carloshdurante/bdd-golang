package truck_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/carloshdurante/bdd-golang/truck"
)

var _ = Describe("Truck", func() {

	Describe("GetTruck", func() {
		TruckValid := truck.Truck{
			Brand: "Ford",
			Model: "F150",
		}
		Context("when the truck is valid", func() {
			It("should return the truck", func() {
				Expect(TruckValid.GetTruck(TruckValid)).Should(Equal("Brand: Ford\nModel: F150"))
			})
		})

		Context("when the truck is empty", func() {
			TruckNotValid := truck.Truck{}
			It("should return truck is empty", func() {
				Expect(TruckNotValid.GetTruck(TruckNotValid)).Should(Equal("Truck is empty"))
			})
		})
	})
})
