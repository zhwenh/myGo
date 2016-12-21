package main

import (
	"fmt"
)

func MultiplyArray(arr1 []float64, num float64) {
	for i := range arr1 {
		arr1[i] *= num
	}
}


func main() {
	arrf := []float64{4,5,6, 7,8,9, 10}
	MultiplyArray(arrf, 1000)
	fmt.Println(arrf)
	var arr []int
	arr = make([]int, 1, 100)
	arr[0] = 9
	fmt.Println(arr[0])
	fmt.Println(len(arr), cap(arr))
	arr = append(arr, 23)
	fmt.Println(len(arr), cap(arr))
	arr2 := arr[3:15]
	fmt.Println(arr2, len(arr2), cap(arr2))
	arr3 := []int{4,5,6, 7,8,9, 10}
	fmt.Println(arr3, len(arr3), cap(arr3))
	arr4 := arr3[4:6]
	fmt.Println(arr4, len(arr4), cap(arr4))

}
