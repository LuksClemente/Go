package main

import "fmt"

func recursao(v []int, n int){
	if(n < 2){
		return
	}
	v1 := make([]int, n/2)
	v2 := make([]int, n/2)
	for i := 0; i < n/2; i = i + 1{
		v[i]++;
	}
	recursao(v1, n/2)
	recursao(v2, n/2)
}
func main(){
	v := make([]int, 10)
	v = []int{2, 3, 5, 4, 1, 7, 8, 6, 9, 0}
	recursao(v, 10)
	for i := 0; i < 10; i++{
		fmt.Printf("%d -> ", v[i])
	}
}