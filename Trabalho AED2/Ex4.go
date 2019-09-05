package main

import "fmt"
import "math"

func recursao(v []int, n int){
	if(n < 2){
		return
	}
	var t int = int(math.Pow(float64(n), float64(n)))
	
	v1 := make([]int, n/2)
	
	var w int = int(math.Pow(float64(2), float64(n)))
	
	for i := 0; i < t; i++{
		v[i] = i
		fmt.Printf("%d -> ", v[i])
	}	
	for i := 0; i < w; i++{
		recursao(v1, n/2)
	}
}
func main(){
	v := make([]int, 5)
	recursao(v, 2)
}