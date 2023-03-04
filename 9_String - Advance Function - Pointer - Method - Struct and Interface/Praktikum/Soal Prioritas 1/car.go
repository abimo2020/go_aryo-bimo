package main

// "strconv"
// "strings"

type Car struct {
	carType string
	fuelIn  int
}

func (c Car) countDistance() int {
	return c.fuelIn * 2 / 3
}
func main() {
	var car1 = new(Car)
	car1.carType = "Sedan"
	car1.fuelIn = 15
	println("Mobil dengan type", car1.carType, "yang berisi", car1.fuelIn, "L bensin dapat menempuh perjalanan sejauh", car1.countDistance(), "mill")
}
