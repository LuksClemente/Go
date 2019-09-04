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
	for i := 0; i < n/2; i++{
		for j := 0; j < n/2; j++{
			v1[i] = 0
			v2[i] = 0
		}
		fmt.Printf("vetor1: %d, ", v1[i])
		fmt.Printf("vetor2: %d, ", v2[i])
		fmt.Printf("\n") 
	} 
	recursao(v1, n/2)
}
func main(){
	v := make([]int, 10)
	v = []int{5, 8, 6, 7, 3, 1, 0, 2, 9, 4}
	recursao(v, 10)
}