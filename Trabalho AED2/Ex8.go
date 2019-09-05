package main

import "fmt"
import "math"

func recursao(v []int, n int){
	if(n < 4){
		return
	}
	var t float64 = math.Pow(float64(n), float64(0.51))
	
	v1 := make([]int, n/4)
	v2 := make([]int, n/4)
	
	for i := 0; float64(i) < t; i++{
		v[i] = i
		fmt.Printf("%d -> ", v[i])
	}
	recursao(v1, n/4)
	recursao(v2, n/4)
}
func main(){
	v := make([]int, 10)
	recursao(v, 2)
}