package main

import "fmt"

func recursao(v []int, n int){
	if(n < 2){
		return
	}
	v1 := make([]int, n/2) 
	v2 := make([]int, n/2)
	recursao(v1, n/2)
	recursao(v2, n/2)
	var count int = 0
	for i := 0; i < n; i++{
		for j := 0; j < n; j++{
			count++
		}
	}
	fmt.Printf("%d\n", count)
	recursao(v1, n/2)
	recursao(v2, n/2)
}
func main(){
	v := make([]int, 10)
	v = []int{5, 8, 6, 7, 3, 1, 0, 2, 9, 4}
	recursao(v, 10)
}