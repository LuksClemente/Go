package main

import "fmt"

func recursao(v []int, n int){
	if(n < 3){
		return
	}
	
	v1 := make([]int, n/3) 
	v2 := make([]int, n/3)
	v3 := make([]int, n/3)
	var cont int = 0
	for i := 0; i < n; i = i + 2{
		cont++
		v[i] = i
	}
	fmt.Printf("%d\n", cont)
	recursao(v1, n/3)
	recursao(v2, n/3)
	recursao(v3, n/3)
}
func main(){
	v := make([]int, 15)
	recursao(v, 15)
}