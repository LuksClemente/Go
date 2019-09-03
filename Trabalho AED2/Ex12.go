package main

import "fmt"

func recursao(v []int, n int){
	if(n < 2){
		return
	}
	v1 := make([]int, n/2) 
	v2 := make([]int, n/2)
	var cont int = 0
	for i := 0; i < n; i = i + 1{
		cont++;
		v[i] = cont
	}
	recursao(v1, n/2)
	recursao(v2, n/2)
	recursao(v, n-1)
}
func main(){
	v := make([]int, 10)
	recursao(v, 10)
	for i := 0; i < 10; i++{
		fmt.Printf("%d -> ", v[i])
	}
}