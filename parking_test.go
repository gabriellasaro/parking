package parking

import "testing"

func TestParking_TotalSpaces(t *testing.T) {
	p := NewParking(10, 20, 5)

	totalSpaces := p.TotalSpaces()
	expectedTotal := uint32(10 + 20 + 5)

	if totalSpaces != expectedTotal {
		t.Errorf("expected total spaces to be %d, got %d", expectedTotal, totalSpaces)
	}
}

func TestParking_TotalOfParkedVehicles(t *testing.T) {
	p := NewParking(10, 20, 5)

	p.parkedCars = 5
	p.parkedMotorcycles = 8
	p.bigCarsParked = 3

	totalParked := p.TotalOfParkedVehicles()
	expectedTotal := uint32(5 + 8 + 3)

	if totalParked != expectedTotal {
		t.Errorf("expected total parked vehicles to be %d, got %d", expectedTotal, totalParked)
	}
}

func TestParking_AvailableSpaces(t *testing.T) {
	p := NewParking(10, 20, 5)
	p.parkedCars = 3
	p.parkedMotorcycles = 5
	p.bigCarsParked = 1

	availableSpaces := p.AvailableSpaces()
	expectedAvailable := uint32(10 + 20 + 5 - (3 + 5 + 1))

	if availableSpaces != expectedAvailable {
		t.Errorf("expected available spaces to be %d, got %d", expectedAvailable, availableSpaces)
	}
}

func TestParking_AvailableSpacesByVehicle(t *testing.T) {
	p := NewParking(10, 20, 5)
	p.parkedCars = 3
	p.parkedMotorcycles = 5
	p.bigCarsParked = 1

	availableSpacesCar, _ := p.AvailableSpacesByVehicle(VehicleTypeCar)
	expectedAvailableCar := uint32(10 - 3)
	if availableSpacesCar != expectedAvailableCar {
		t.Errorf("expected available spaces for car to be %d, got %d", expectedAvailableCar, availableSpacesCar)
	}

	availableSpacesMotorcycle, _ := p.AvailableSpacesByVehicle(VehicleTypeMotorcycle)
	expectedAvailableMotorcycle := uint32(20 - 5)
	if availableSpacesMotorcycle != expectedAvailableMotorcycle {
		t.Errorf("expected available spaces for motorcycle to be %d, got %d", expectedAvailableMotorcycle, availableSpacesMotorcycle)
	}

	availableSpacesVan, _ := p.AvailableSpacesByVehicle(VehicleTypeVan)
	expectedAvailableVan := uint32(5 - 1)
	if availableSpacesVan != expectedAvailableVan {
		t.Errorf("expected available spaces for van to be %d, got %d", expectedAvailableVan, availableSpacesVan)
	}

	_, err := p.AvailableSpacesByVehicle(9) // Using an invalid vehicle type
	if err == nil {
		t.Errorf("expected error for invalid vehicle type, but got nil")
	}
}

func TestParking_AvailableSpacesByType(t *testing.T) {
	p := NewParking(10, 20, 5)
	p.parkedCars = 3
	p.parkedMotorcycles = 5
	p.bigCarsParked = 1

	availableSpacesCar, _ := p.AvailableSpacesByType(TypeOfSpaceCar)
	expectedAvailableCar := uint32(10 - 3)
	if availableSpacesCar != expectedAvailableCar {
		t.Errorf("expected available spaces for car type to be %d, got %d", expectedAvailableCar, availableSpacesCar)
	}

	availableSpacesMotorcycle, _ := p.AvailableSpacesByType(TypeOfSpaceMotorcycle)
	expectedAvailableMotorcycle := uint32(20 - 5)
	if availableSpacesMotorcycle != expectedAvailableMotorcycle {
		t.Errorf("expected available spaces for motorcycle type to be %d, got %d", expectedAvailableMotorcycle, availableSpacesMotorcycle)
	}

	availableSpacesBigCar, _ := p.AvailableSpacesByType(TypeOfSpaceBigCar)
	expectedAvailableBigCar := uint32(5 - 1)
	if availableSpacesBigCar != expectedAvailableBigCar {
		t.Errorf("expected available spaces for big car type to be %d, got %d", expectedAvailableBigCar, availableSpacesBigCar)
	}

	_, err := p.AvailableSpacesByType(9) // Using an invalid space type
	if err == nil {
		t.Errorf("expected error for invalid space type, but got nil")
	}
}

func TestParking_SpacesOccupiedByVehicle(t *testing.T) {
	p := NewParking(10, 20, 0)

	car1, _ := NewCar()
	car2, _ := NewCar()
	motorcycle, _ := NewMotorBike()
	van, _ := NewVan()

	_ = p.Park(car1)
	_ = p.Park(car2)
	_ = p.Park(motorcycle)
	_ = p.Park(van)

	spacesOccupiedByCars := p.SpacesOccupiedByVehicle(VehicleTypeCar)
	expectedSpacesOccupiedByCars := car1.OccupiedSpaces() + car2.OccupiedSpaces()
	if spacesOccupiedByCars != expectedSpacesOccupiedByCars {
		t.Errorf("expected spaces occupied by cars to be %d, got %d", expectedSpacesOccupiedByCars, spacesOccupiedByCars)
	}

	spacesOccupiedByMotorcycles := p.SpacesOccupiedByVehicle(VehicleTypeMotorcycle)
	expectedSpacesOccupiedByMotorcycles := motorcycle.OccupiedSpaces()
	if spacesOccupiedByMotorcycles != expectedSpacesOccupiedByMotorcycles {
		t.Errorf("expected spaces occupied by motorcycles to be %d, got %d", expectedSpacesOccupiedByMotorcycles, spacesOccupiedByMotorcycles)
	}

	spacesOccupiedByVans := p.SpacesOccupiedByVehicle(VehicleTypeVan)
	expectedSpacesOccupiedByVans := van.OccupiedSpaces()
	if spacesOccupiedByVans != expectedSpacesOccupiedByVans {
		t.Errorf("expected spaces occupied by vans to be %d, got %d", expectedSpacesOccupiedByVans, spacesOccupiedByVans)
	}
}

func TestParking_IsEmpty(t *testing.T) {
	p := NewParking(10, 20, 5)

	isEmpty := p.IsEmpty()
	if !isEmpty {
		t.Errorf("expected parking to be empty, but it's not")
	}

	car, _ := NewCar()
	_ = p.Park(car)

	isEmpty = p.IsEmpty()
	if isEmpty {
		t.Errorf("expected parking to not be empty, but it is")
	}
}

func TestParking_NoSpaceAvailable(t *testing.T) {
	p := NewParking(2, 1, 1)

	noSpaceAvailable := p.NoSpaceAvailable()
	if noSpaceAvailable {
		t.Errorf("expected there to be space available, but NoSpaceAvailable returned true")
	}

	car1, _ := NewCar()
	car2, _ := NewCar()
	motorcycle, _ := NewMotorBike()
	van, _ := NewVan()

	_ = p.Park(car1)
	_ = p.Park(car2)
	_ = p.Park(motorcycle)
	_ = p.Park(van)

	noSpaceAvailable = p.NoSpaceAvailable()
	if !noSpaceAvailable {
		t.Errorf("expected no space available, but NoSpaceAvailable returned false")
	}
}

func TestParking_BestSpaceAvailable(t *testing.T) {
	p := NewParking(10, 20, 5)

	car, err := NewCar()
	if err != nil {
		t.Fatalf("error creating car: %v", err)
	}

	van, err := NewVan()
	if err != nil {
		t.Fatalf("error creating van: %v", err)
	}

	bestSpaceCar, err := p.BestSpaceAvailable(car)
	if err != nil {
		t.Errorf("error getting best space for car: %v", err)
	}

	if bestSpaceCar != TypeOfSpaceCar {
		t.Errorf("expected best space for car to be TypeOfSpaceCar, got %v", bestSpaceCar)
	}

	bestSpaceVan, err := p.BestSpaceAvailable(van)
	if err != nil {
		t.Errorf("error getting best space for van: %v", err)
	}

	if bestSpaceVan != TypeOfSpaceBigCar {
		t.Errorf("expected best space for van to be TypeOfSpaceBigCar, got %v", bestSpaceVan)
	}
}

func TestParking_Park(t *testing.T) {
	p := NewParking(1, 2, 0)

	car, _ := NewCar()
	motorcycle, _ := NewMotorBike()
	van, _ := NewVan()

	err := p.Park(car)
	if err != nil {
		t.Errorf("error parking car: %v", err)
	}

	if p.TotalOfParkedVehicles() != 1 {
		t.Errorf("expected 1 parked vehicle, got %d", p.TotalOfParkedVehicles())
	}

	err = p.Park(motorcycle)
	if err != nil {
		t.Errorf("error parking motorcycle: %v", err)
	}

	if p.TotalOfParkedVehicles() != 2 {
		t.Errorf("expected 2 parked vehicles, got %d", p.TotalOfParkedVehicles())
	}

	err = p.Park(van)
	if err == nil {
		t.Errorf("expected error parking van, but got nil")
	}

	if p.TotalOfParkedVehicles() != 2 {
		t.Errorf("expected 2 parked vehicles, got %d", p.TotalOfParkedVehicles())
	}
}

func TestSpaceTypeByVehicle(t *testing.T) {
	carSpace, _ := spaceTypeByVehicle(VehicleTypeCar)
	if carSpace != TypeOfSpaceCar {
		t.Errorf("expected space type for car to be TypeOfSpaceCar, got %v", carSpace)
	}

	motorcycleSpace, _ := spaceTypeByVehicle(VehicleTypeMotorcycle)
	if motorcycleSpace != TypeOfSpaceMotorcycle {
		t.Errorf("expected space type for motorcycle to be TypeOfSpaceMotorcycle, got %v", motorcycleSpace)
	}

	vanSpace, _ := spaceTypeByVehicle(VehicleTypeVan)
	if vanSpace != TypeOfSpaceBigCar {
		t.Errorf("expected space type for van to be TypeOfSpaceBigCar, got %v", vanSpace)
	}

	_, err := spaceTypeByVehicle(9) // Using an invalid vehicle type
	if err == nil {
		t.Errorf("expected error for invalid vehicle type, but got nil")
	}
}
