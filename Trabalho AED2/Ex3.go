package main

import "fmt"
import "math"

func recursao(v []int, n int){
	if(n < 2){
		return
	}
	var nf float64 = float64(n)
	var t int = int(math.Pow(2, nf))
	V := make([]int, t) 
	v1 := make([]int, n/2)
	for i := 0; i < t; i++{
		V[i] = i;
		fmt.Printf("%d ->", V[i])
	}
	recursao(v1, n)
}
func main(){
	v := make([]int, 4)
	var n int = 4
	recursao(v, n)
}
