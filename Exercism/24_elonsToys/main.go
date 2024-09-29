package main

import "fmt"

/*A method is a function with a special receiver argument.
The receiver appears in its own argument list between func keyword and the name of the method.
func (receiver type) MethodName(parameters) (returnTypes) {}
You can only define a method with a receiver whose type is defined in the same package as the method.
*/

// FROM PREVIOUS EXERCISE
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

func NewCar(speed int, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

/*
The *Car indicates that this method has a pointer receiver to a type named Car.
A pointer receiver allows the method to modify the original Car struct, rather than a copy of it.
In this case, car is a pointer to a Car instance.
*/
func (car *Car) Drive() {
	if car.battery-car.batteryDrain >= 0 {
		car.distance += car.speed
		car.battery -= car.batteryDrain
		fmt.Println(car)
	}
}

func (car *Car) DisplayDistance() string {
	displayMessage := fmt.Sprintf("Driven %d meters", car.distance)
	return displayMessage
}

func (car *Car) DisplayBattery() string {
	displayMessage := fmt.Sprintf("Battery at %d%%", car.battery)
	return displayMessage
}

func (car *Car) CanFinish(trackDistance int) bool {
	car.distance = (car.battery / car.batteryDrain) * car.speed
	if car.distance >= trackDistance {
		fmt.Println("true")
		return true
	} else {
		return false
	}
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	car.Drive()
	fmt.Println(car.DisplayDistance())
	fmt.Println(car.DisplayBattery())
	trackDistance := 100
	car.CanFinish(trackDistance)
}
