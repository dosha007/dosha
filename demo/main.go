package main

import (
	"fmt"
	"transport_workspace/car"
	"transport_workspace/landvehicle"
	"transport_workspace/vehicle"
)

// Функция для демонстрации полиморфизма
func demonstrateTransport(transport vehicle.Transport) {
	fmt.Println("\nДемонстрация работы транспорта через интерфейс:")
	transport.DisplayInfo()
	transport.Move()
	transport.CarryPassengers(2)
	transport.Stop()
}

func main() {
	fmt.Println("Демонстрация иерархии транспортных средств")

	// Создаем транспортное средство
	ship := vehicle.NewVehicle("Титаник", 50.0, 3547, 52310000)
	fmt.Println("\n1. Базовое транспортное средство:")
	ship.DisplayInfo()
	ship.Move()
	ship.CarryPassengers(2000)
	ship.Stop()

	// Создаем наземный транспорт
	train := landvehicle.NewLandVehicle("Скоростной поезд", 350.0, 500, 400000, 48, "электричество")
	fmt.Println("\n2. Наземный транспорт:")
	train.DisplayLandInfo()
	train.DriveOnRoad()
	train.Refuel(10000)
	train.ChangeWheels(50)
	train.Brake()

	// Создаем автомобиль
	myCar := car.NewCar("Toyota", "Camry", 2022, 220.0, 5, 1560, 4, "бензин")
	fmt.Println("\n3. Автомобиль:")
	myCar.DisplayCarInfo()
	myCar.Honk()
	myCar.Accelerate(100)
	myCar.Accelerate(250) // Превысит максимальную скорость
	myCar.Park()
	myCar.Service()

	// Демонстрация работы с интерфейсом Transport
	demonstrateTransport(ship)
	demonstrateTransport(&train.Vehicle)
	demonstrateTransport(&myCar.Vehicle)
}
