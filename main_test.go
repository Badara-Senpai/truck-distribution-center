package main

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestProcessTruck(t *testing.T) {
	t.Run("should load and unload a truck cargo", func(t *testing.T) {
		nt := &NormalTruck{id: "Truck-1", cargo: 42}
		et := &ElectricTruck{id: "e-Truck-1"}

		if err := processTruck(nt); err != nil {
			t.Fatalf("Error processing truck: %s", err)
		}

		if err := processTruck(et); err != nil {
			t.Fatalf("Error processing truck: %s", err)
		}

		// asserting
		if nt.cargo != 0 {
			t.Fatal("Normal truck cargo should be 0")
		}

		if et.battery != -2 {
			t.Fatal("Electric truck battery should be -2")
		}
	})
}
