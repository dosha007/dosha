package car

import (
	"fmt"
    "transport_workspace/landvehicle"
)

type Car struct {
    landvehicle.LandVehicle
    Brand string
    Model string
    Year  int
}

// ... (остальной код Car)
