package main

type ParkingSystem struct {
	num1 int
	num2 int
	num3 int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	var ps ParkingSystem
	ps.num1 = big
	ps.num2 = medium
	ps.num3 = small
	return ps
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.num1 > 0 {
			this.num1--
			return true
		}
		return false
	case 2:
		if this.num2 > 0 {
			this.num2--
			return true
		}
		return false
	case 3:
		if this.num3 > 0 {
			this.num3--
			return true
		}
		return false
	default:
		return false
	}
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
