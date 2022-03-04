package truck_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTruck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Truck Suite")
}
