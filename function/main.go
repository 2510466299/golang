package main

import (
	"fmt"
	"golang/pactangle"
)

/*
func funcname(parametername type)return type{

}
*/

func main() {
	price, no := 10, 20
	totalPrice := pactangle.Caculatebill(price, no)
	fmt.Println("The total price is:", totalPrice)

	lenth, width := 4.6, 13.5
	area, perimeter := pactangle.AreaAndPerimeter(lenth, width)
	fmt.Printf("The area is:%.2f,The perimeter is:%.2f\n", area, perimeter)
}
