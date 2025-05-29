package vehicle

import "fmt"

type Transport interface {
    Move()
    Stop()
    CarryPassengers(num int)
    SetMaxSpeed(speed float64)
    DisplayInfo()
}

type Vehicle struct {
    Name     string
    MaxSpeed float64
    Capacity int
    Weight   float64
}

func NewVehicle(name string, maxSpeed float64, capacity int, weight float64) *Vehicle {
    return &Vehicle{
        Name:     name,
        MaxSpeed: maxSpeed,
        Capacity: capacity,
        Weight:   weight,
    }
}

// ... (остальные методы Vehicle)
