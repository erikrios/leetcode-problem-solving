package main

import "fmt"

func main() {
	parkingSystem := Constructor(1, 1, 0)
	fmt.Println(parkingSystem.AddCar(1))
	fmt.Println(parkingSystem.AddCar(2))
	fmt.Println(parkingSystem.AddCar(3))
	fmt.Println(parkingSystem.AddCar(1))
}

type ParkingSystem struct {
	Big    int
	Medium int
	Small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Big:    big,
		Medium: medium,
		Small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.Big > 0 {
			this.Big--
			return true
		}
		return false
	case 2:
		if this.Medium > 0 {
			this.Medium--
			return true
		}
		return false
	default:
		if this.Small > 0 {
			this.Small--
			return true
		}
		return false
	}
}
