module transport_workspace/car

go 1.24.3

require (
    transport_workspace/landvehicle v0.0.0
    transport_workspace/vehicle v0.0.0
)

replace (
    transport_workspace/landvehicle => ../landvehicle
    transport_workspace/vehicle => ../vehicle
)
