package main

import "math"

func recursao(v []int, n int){
	if(n < 3){
		return
	}
	v1 := make([]int, n/3) 
	v2 := make([]int, n/3)
	v3 := make([]int, n/3)
	var i int = 0
	
	for ok := true; ok; ok = (i < int(math.Sqrt(float64(n)))){
		v1[i] = 0
		v2[i] = 0
		v3[i] = 0
		i++
	}
	
	recursao(v1, n/3)
	recursao(v2, n/3)
	recursao(v3, n/3)
}
func main(){
	v := make([]int, 15)
	v = []int{5, 6, 2, 9, 11, 0, 4, 18, 14, 21, 7, 3, 25, 26, 19}
	recursao(v, 3)
}