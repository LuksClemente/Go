package main

import "fmt"
import "math"

func recursao(v []int, n int){
	if(n < 2){
		return
	}
	v1 := make([]int, n/2) 
	v2 := make([]int, n/2)
	recursao(v1, n/2)
	recursao(v2, n/2)
	for i := 0; float64(i) <= float64(n)/math.Log2(float64(n)); i++{
		fmt.Printf("%d\n", i)
	}
	recursao(v1, n/2)
	recursao(v2, n/2)
}
func main(){
	v := make([]int, 15)
	v = []int{5, 6, 2, 9, 11, 0, 4, 18, 14, 21, 7, 3, 25, 26, 19}
	recursao(v, 15)
}