package main

import "fmt"

/*func main() {
	fatia := make([]float64, 5)
	fmt.Println(fatia)
}*/

/*func main() {

	arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	fatia := arr[2:5]
	fmt.Println(fatia)
}*/

/*func main() {
	fatia1 := []int{1, 2, 3}
	fatia2 := append(fatia1, 2, 4, 5)
	fmt.Println(fatia1, fatia2)
}*/

func main() {
	fatia1 := []int{1, 2, 3}
	fatia2 := make([]int, 1)
	copy(fatia2, fatia1)
	fmt.Println(fatia1, fatia2)
}
