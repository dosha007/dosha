package landvehicle

import (
	"fmt"
    "transport_workspace/vehicle"
)

type LandVehicle struct {
    vehicle.Vehicle
    WheelCount int
    FuelType   string
}

// ... (остальной код LandVehicle)
