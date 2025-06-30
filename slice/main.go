package main

import "fmt"

func main() {
	/*
		var arr = [...]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
		slice := arr[1:4]
		fmt.Printf("slice的类型:%T\n", slice)
		for index, value := range slice {
			fmt.Println(index, value)
		}
		slice_make := make([]int, 5, 10)
		//slice is need to make'{}' to use,if not ,it will be nil .
		fmt.Println(slice_make)
		fmt.Printf("slice_make的类型:%T\n", slice_make)
		slice_make[0] = 100
		slice_make[1] = 200
		slice_make[2] = 300
		slice_make[3] = 400
		slice_make[4] = 500
		for i := 0; i < 6; i++ {
			slice_make = append(slice_make, i)
			fmt.Println(cap(slice_make))
		}
		var slice_nil []int
		if slice_nil == nil {
			fmt.Println("slice_nil是nil")
		}
		slice_make_nil := make([]int, 0)
		if slice_make_nil == nil {
			fmt.Println("slice_make_nil是nil")
		}
	*/

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5, 5)
	copy(b, a)
	fmt.Println(a)
	fmt.Println(b)
	c := b
	fmt.Printf("a's address:%p\n", a)
	fmt.Printf("b's address:%p\n", b)
	fmt.Printf("c's address:%p\n", c)
	//slice'delete
	slice_delete := append(a[0:2], a[3:]...)
	fmt.Println(slice_delete)
	//slice'delete is append (slice[0:i],slice[i+1:]...)
	//slice'insert
	slice_insert := append(a[:2], append([]int{100}, a[2:]...)...)
	fmt.Println(slice_insert)
	//slice'update
	slice_update := append(a[:2], append([]int{200}, a[2:]...)...)
	fmt.Println(slice_update)
	x := make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		x = append(x, fmt.Sprintf("%v", i))
	}

}
