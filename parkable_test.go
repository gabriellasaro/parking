package parking

import (
	"testing"
)

func TestCar(t *testing.T) {
	car, err := NewCar()
	if err != nil {
		t.Fatalf("error creating car: %v", err)
	}

	if car.OccupiedSpaces() != 1 {
		t.Errorf("expected occupied spaces for car to be 1, got %d", car.OccupiedSpaces())
	}

	if !car.SpaceAllowed(TypeOfSpaceCar) {
		t.Error("expected car to be allowed in TypeOfSpaceCar")
	}

	if car.SpaceAllowed(TypeOfSpaceMotorcycle) {
		t.Error("expected car to not be allowed in TypeOfSpaceMotorcycle")
	}

	err = car.Park(TypeOfSpaceBigCar)
	if err != nil {
		t.Errorf("error parking car: %v", err)
	}

	if car.TypeOfSpace() != TypeOfSpaceBigCar {
		t.Errorf("expected car to be parked in TypeOfSpaceBigCar, got %v", car.TypeOfSpace())
	}
}

func TestVan(t *testing.T) {
	van, err := NewVan()
	if err != nil {
		t.Fatalf("error creating van: %v", err)
	}

	if van.OccupiedSpaces() != 1 {
		t.Errorf("expected occupied spaces for van to be 1, got %d", van.OccupiedSpaces())
	}

	if !van.SpaceAllowed(TypeOfSpaceCar) {
		t.Error("expected van to be allowed in TypeOfSpaceCar")
	}

	if van.SpaceAllowed(TypeOfSpaceMotorcycle) {
		t.Error("expected van to not be allowed in TypeOfSpaceMotorcycle")
	}

	err = van.Park(TypeOfSpaceBigCar)
	if err != nil {
		t.Errorf("error parking van: %v", err)
	}

	if van.TypeOfSpace() != TypeOfSpaceBigCar {
		t.Errorf("expected van to be parked in TypeOfSpaceBigCar, got %v", van.TypeOfSpace())
	}
}

func TestParkable_Plate(t *testing.T) {
	motorcycle, err := NewMotorBike()
	if err != nil {
		t.Fatalf("error creating motorcycle: %v", err)
	}

	plate := motorcycle.Plate()
	if len(plate) != 36 {
		t.Errorf("expected plate length to be 36, got %d", len(plate))
	}
}

func TestParkable_Vehicle(t *testing.T) {
	motorcycle, err := NewMotorBike()
	if err != nil {
		t.Fatalf("error creating motorcycle: %v", err)
	}

	vehicle := motorcycle.Vehicle()
	if vehicle != VehicleTypeMotorcycle {
		t.Errorf("expected vehicle type to be VehicleTypeMotorcycle, got %v", vehicle)
	}
}

func TestParkable_Park(t *testing.T) {
	motorcycle, err := NewMotorBike()
	if err != nil {
		t.Fatalf("error creating motorcycle: %v", err)
	}

	err = motorcycle.Park(TypeOfSpaceCar)
	if err != nil {
		t.Errorf("error parking motorcycle: %v", err)
	}

	if motorcycle.TypeOfSpace() != TypeOfSpaceCar {
		t.Errorf("expected motorcycle to be parked in TypeOfSpaceCar, got %v", motorcycle.TypeOfSpace())
	}
}

func TestParkable_TypeOfSpace(t *testing.T) {
	motorcycle, err := NewMotorBike()
	if err != nil {
		t.Fatalf("error creating motorcycle: %v", err)
	}

	space := motorcycle.TypeOfSpace()
	if space != TypeOfSpaceNotParked {
		t.Errorf("expected type of space to be TypeOfSpaceNotParked, got %v", space)
	}

	if err := motorcycle.Park(TypeOfSpaceMotorcycle); err != nil {
		t.Fatalf("error parking motorcycle: %v", err)
	}

	space = motorcycle.TypeOfSpace()
	if space != TypeOfSpaceMotorcycle {
		t.Errorf("expected type of space to be TypeOfSpaceMotorcycle, got %v", space)
	}
}

func TestParkable_SpaceAllowed(t *testing.T) {
	car, err := NewCar()
	if err != nil {
		t.Fatalf("error creating car: %v", err)
	}

	if !car.SpaceAllowed(TypeOfSpaceCar) {
		t.Error("expected car to be allowed in TypeOfSpaceCar")
	}

	if car.SpaceAllowed(TypeOfSpaceMotorcycle) {
		t.Error("expected car to not be allowed in TypeOfSpaceMotorcycle")
	}
}

func TestParkableOccupiedSpaces(t *testing.T) {
	van, err := NewVan()
	if err != nil {
		t.Fatalf("error creating car: %v", err)
	}

	if err := van.Park(TypeOfSpaceCar); err != nil {
		t.Fatalf("error parking van: %v", err)
	}

	occupiedSpaces := van.OccupiedSpaces()
	if occupiedSpaces != 3 {
		t.Errorf("expected occupied spaces for van (type of space for car) to be 3, got %d", occupiedSpaces)
	}
}

func TestGeneratePlate(t *testing.T) {
	plate, err := generatePlate()
	if err != nil {
		t.Fatalf("error generating plate: %v", err)
	}

	if len(plate) != 36 {
		t.Errorf("expected plate length to be 36, got %d", len(plate))
	}
}
