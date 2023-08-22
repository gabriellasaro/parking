package parking

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrTypeOfSpaceNotAllowed = errors.New("type of space not allowed")
	ErrInvalidVehicleType    = errors.New("invalid vehicle type")
	ErrInvalidSpaceType      = errors.New("invalid space type")

	listSpaces = []TypeOfSpace{
		TypeOfSpaceCar, TypeOfSpaceBigCar, TypeOfSpaceMotorcycle,
	}
)

type TypeOfSpace uint8

const (
	TypeOfSpaceNotParked TypeOfSpace = iota
	TypeOfSpaceCar
	TypeOfSpaceMotorcycle
	TypeOfSpaceBigCar
)

type VehicleType uint8

const (
	VehicleTypeCar VehicleType = iota
	VehicleTypeMotorcycle
	VehicleTypeVan
)

type Parkable interface {
	Plate() string
	Vehicle() VehicleType
	TypeOfSpace() TypeOfSpace
	Park(space TypeOfSpace) error
	OccupiedSpaces() uint32
	SpaceAllowed(space TypeOfSpace) bool
}

type parkable struct {
	plate       string
	vehicle     VehicleType
	typeOfSpace TypeOfSpace
}

func (p *parkable) Plate() string {
	return p.plate
}

func (p *parkable) Vehicle() VehicleType {
	return p.vehicle
}

func (p *parkable) TypeOfSpace() TypeOfSpace {
	return p.typeOfSpace
}

func (p *parkable) SpaceAllowed(space TypeOfSpace) bool {
	switch space {
	case TypeOfSpaceCar, TypeOfSpaceMotorcycle, TypeOfSpaceBigCar:
		return true
	default:
		return false
	}
}

func (p *parkable) Park(space TypeOfSpace) error {
	if p.SpaceAllowed(space) {
		p.typeOfSpace = space

		return nil
	}

	return ErrTypeOfSpaceNotAllowed
}

func (p *parkable) OccupiedSpaces() uint32 {
	return 1
}

type Car struct {
	parkable
}

func (c *Car) SpaceAllowed(space TypeOfSpace) bool {
	switch space {
	case TypeOfSpaceCar, TypeOfSpaceBigCar:
		return true
	default:
		return false
	}
}

func (c *Car) Park(space TypeOfSpace) error {
	if c.SpaceAllowed(space) {
		c.typeOfSpace = space

		return nil
	}

	return fmt.Errorf("car: %w", ErrTypeOfSpaceNotAllowed)
}

type Van struct {
	parkable
}

func (v *Van) SpaceAllowed(space TypeOfSpace) bool {
	switch space {
	case TypeOfSpaceCar, TypeOfSpaceBigCar:
		return true
	default:
		return false
	}
}

func (v *Van) Park(space TypeOfSpace) error {
	if v.SpaceAllowed(space) {
		v.typeOfSpace = space

		return nil
	}

	return fmt.Errorf("van: %w", ErrTypeOfSpaceNotAllowed)
}

func (v *Van) OccupiedSpaces() uint32 {
	if v.typeOfSpace == TypeOfSpaceCar {
		return 3
	}

	return 1
}

type Motorcycle struct {
	parkable
}

func NewCar() (*Car, error) {
	plate, err := generatePlate()
	if err != nil {
		return nil, err
	}

	return &Car{
		parkable: parkable{
			plate:   plate,
			vehicle: VehicleTypeCar,
		},
	}, nil
}

func NewMotorBike() (*Motorcycle, error) {
	plate, err := generatePlate()
	if err != nil {
		return nil, err
	}

	return &Motorcycle{
		parkable: parkable{
			plate:   plate,
			vehicle: VehicleTypeMotorcycle,
		},
	}, nil
}

func NewVan() (*Van, error) {
	plate, err := generatePlate()
	if err != nil {
		return nil, err
	}

	return &Van{
		parkable: parkable{
			plate:   plate,
			vehicle: VehicleTypeVan,
		},
	}, nil
}

func generatePlate() (string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}
