package car

const (
	PointValue = iota
	PointBicycleValue
	PointCarValue
)

var cars = "🚗🚕🚙🏎🚓🚑🚒🚌🚎🚐🚚🚛🚜🦼🛴🚲🛵🏍🦽🛺"

type PointCar struct {
	v int
}

func NewPointCar(v int) PointCar {
	return PointCar{v}
}

func (p PointCar) Render() string {
	return string(cars[p.v])
}
func (p PointCar) Value() int {
	return PointCarValue
}
type PointBicycle struct {
}

func (p PointBicycle) Render() string {
	return "🚴"
}
func (p PointBicycle) Value() int {
	return PointBicycleValue
}
