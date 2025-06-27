package pactangle

func AreaAndPerimeter(lenth, width float64) (area, perimeter float64) {
	area = lenth * width
	perimeter = 2 * (lenth + width)
	return
}
func Caculatebill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}
