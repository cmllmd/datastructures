package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 100
	arr[1] = 200
	arr[2] = 300
	arr[3] = 400
	arr[4] = 500

	arr2 := [5]int{1, 2, 3, 4, 5} //definindo e populando o array

	fmt.Println(arr, arr2)
}
