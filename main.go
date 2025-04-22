package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

var (
	ErrTruckNotFound  = errors.New("truck not found")
	ErrNotImplemented = errors.New("not implemented")
)

func (nt *NormalTruck) LoadCargo() error {
	nt.cargo += 1
	return nil
}

func (nt *NormalTruck) UnloadCargo() error {
	nt.cargo = 0
	return nil
}

func (et *ElectricTruck) LoadCargo() error {
	et.cargo += 1
	et.battery += -1
	return nil
}

func (et *ElectricTruck) UnloadCargo() error {
	et.cargo = 0
	et.battery += -1
	return nil
}

// processTruck handles the loading and the unloading of a truck
func processTruck(truck Truck) error {
	fmt.Printf("processing truck %v\n", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	return nil
}

func main() {
	nt := &NormalTruck{id: "Truck-1"}
	et := &ElectricTruck{id: "e-Truck-1"}

	if err := processTruck(nt); err != nil {
		log.Fatalf("Error processing truck: %s", err)
	}

	if err := processTruck(et); err != nil {
		log.Fatalf("Error processing truck: %s", err)
	}

	log.Println(nt.cargo)
	log.Println(et.battery)
}
