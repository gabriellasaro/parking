package parking

import (
	"errors"
	"fmt"
)

var (
	ErrNoSpaceAvailable = errors.New("no space available")
)

type Parking struct {
	carSpaces        uint32
	motorcycleSpaces uint32
	bigCarSpaces     uint32

	parkedMotorcycles uint32
	parkedCars        uint32
	bigCarsParked     uint32

	parked []Parkable
}

func NewParking(carSpaces, motorcycleSpaces, bigCarSpaces uint32) Parking {
	return Parking{
		carSpaces:        carSpaces,
		motorcycleSpaces: motorcycleSpaces,
		bigCarSpaces:     bigCarSpaces,
	}
}

func (p *Parking) TotalSpaces() uint32 {
	return p.carSpaces + p.motorcycleSpaces + p.bigCarSpaces
}

func (p *Parking) TotalOfParkedVehicles() uint32 {
	return p.parkedCars + p.parkedMotorcycles + p.bigCarsParked
}

func (p *Parking) AvailableSpaces() uint32 {
	return p.TotalSpaces() - p.TotalOfParkedVehicles()
}

func (p *Parking) AvailableSpacesByVehicle(vehicle VehicleType) (uint32, error) {
	switch vehicle {
	case VehicleTypeCar:
		return p.carSpaces - p.parkedCars, nil
	case VehicleTypeVan:
		return p.bigCarSpaces - p.bigCarsParked, nil
	case VehicleTypeMotorcycle:
		return p.motorcycleSpaces - p.parkedMotorcycles, nil
	}

	return 0, fmt.Errorf("%w: %d", ErrInvalidVehicleType, vehicle)
}

func (p *Parking) AvailableSpacesByType(space TypeOfSpace) (uint32, error) {
	switch space {
	case TypeOfSpaceCar:
		return p.carSpaces - p.parkedCars, nil
	case TypeOfSpaceBigCar:
		return p.bigCarSpaces - p.bigCarsParked, nil
	case TypeOfSpaceMotorcycle:
		return p.motorcycleSpaces - p.parkedMotorcycles, nil
	}

	return 0, fmt.Errorf("%w: %d", ErrInvalidSpaceType, space)
}

func (p *Parking) SpacesOccupiedByVehicle(vehicle VehicleType) uint32 {
	var spaces uint32

	for i := range p.parked {
		if p.parked[i].Vehicle() == vehicle {
			spaces += p.parked[i].OccupiedSpaces()
		}
	}

	return spaces
}

func (p *Parking) IsEmpty() bool {
	return len(p.parked) == 0
}

func (p *Parking) NoSpaceAvailable() bool {
	return p.TotalSpaces() == p.TotalOfParkedVehicles()
}

func (p *Parking) BestSpaceAvailable(parkable Parkable) (TypeOfSpace, error) {
	available, err := p.AvailableSpacesByVehicle(parkable.Vehicle())
	if err != nil {
		return TypeOfSpaceNotParked, err
	}

	if available > 0 {
		return spaceTypeByVehicle(parkable.Vehicle())
	}

	for i := range listSpaces {
		if parkable.SpaceAllowed(listSpaces[i]) {
			available, err := p.AvailableSpacesByType(listSpaces[i])
			if err != nil {
				return TypeOfSpaceNotParked, err
			}

			if available > 0 {
				return listSpaces[i], nil
			}
		}
	}

	return TypeOfSpaceNotParked, ErrNoSpaceAvailable
}

func (p *Parking) Park(parkable Parkable) error {
	if p.NoSpaceAvailable() {
		return ErrNoSpaceAvailable
	}

	space, err := p.BestSpaceAvailable(parkable)
	if err != nil {
		return err
	}

	if err := parkable.Park(space); err != nil {
		return err
	}

	p.addNewParked(parkable)

	return nil
}

func (p *Parking) addNewParked(parkable Parkable) {
	p.parked = append(p.parked, parkable)

	switch parkable.TypeOfSpace() {
	case TypeOfSpaceCar:
		p.parkedCars += parkable.OccupiedSpaces()
	case TypeOfSpaceBigCar:
		p.bigCarsParked += parkable.OccupiedSpaces()
	case TypeOfSpaceMotorcycle:
		p.parkedMotorcycles += parkable.OccupiedSpaces()
	}
}

func spaceTypeByVehicle(vehicle VehicleType) (TypeOfSpace, error) {
	switch vehicle {
	case VehicleTypeCar:
		return TypeOfSpaceCar, nil
	case VehicleTypeVan:
		return TypeOfSpaceBigCar, nil
	case VehicleTypeMotorcycle:
		return TypeOfSpaceMotorcycle, nil
	default:
		return TypeOfSpaceNotParked, fmt.Errorf("%w: %d", ErrInvalidVehicleType, vehicle)
	}
}
