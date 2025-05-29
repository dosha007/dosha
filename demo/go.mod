module transport_workspace/demo

go 1.24.3

require (
    transport_workspace/car v0.0.0
    transport_workspace/landvehicle v0.0.0
    transport_workspace/vehicle v0.0.0
)

replace (
    transport_workspace/car => ../car
    transport_workspace/landvehicle => ../landvehicle
    transport_workspace/vehicle => ../vehicle
)
