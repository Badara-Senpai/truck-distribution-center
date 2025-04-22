package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
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

type contextKey string

var UserIDKey contextKey = "userID"

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
func processTruck(ctx context.Context, truck Truck) error {
	fmt.Printf("Started processing truck %+v \n", truck)

	//access the userID
	//userID := ctx.Value(UserIDKey)

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	delay := time.Second * 1
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		break
	}

	// Simulate processing time
	time.Sleep(time.Second)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	fmt.Printf("Finished processing truck %+v \n", truck)
	return nil
}

func processFleet(ctx context.Context, trucks []Truck) error {
	var wg sync.WaitGroup

	errorsChan := make(chan error, len(trucks))

	for _, t := range trucks {
		wg.Add(1)

		go func(t Truck) {
			if err := processTruck(ctx, t); err != nil {
				log.Println(err)
				errorsChan <- err
			}

			wg.Done()
		}(t)
	}

	wg.Wait()
	close(errorsChan)

	var errs []error
	for err := range errorsChan {
		log.Printf("Error processing truc: %v \n", err)
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("fleet processing had %d errors", len(errs))
	}

	return nil
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, UserIDKey, 42)

	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
	}

	if err := processFleet(ctx, fleet); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	fmt.Println("All trucks processed successfully")
}
