package main

import "fmt"
import "math"

func recursao(v []int, n int){
	if(n < 2){
		return
	}
	v1 := make([]int, n/2)
	for i := 0; i < int(math.Log(float64(n))); i++{
		v[i]++
	}
	for i := 0; float64(i) < math.Sqrt(float64(2)); i++{
		recursao(v1, n/2)
	}
}
func main(){
	v := make([]int, 8)
	v = []int{1, 5, 2, 6, 8, 3, 2, 1}
	recursao(v, 8)
	fmt.Printf("%d", v)
}