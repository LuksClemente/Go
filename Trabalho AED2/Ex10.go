package main

import "fmt"

func recursao(v []int, n int){
	if(n < 4){
		return
	}
	v1 := make([]int, n/4) 
	v2 := make([]int, n/4)
	
	recursao(v1, n/4)
	recursao(v1, n/4)
	recursao(v1, n/4)
	recursao(v1, n/4)
	recursao(v1, n/4)
	recursao(v1, n/4)
	recursao(v1, n/4)
	recursao(v1, n/4)
	
	var fat int = n
	var i int = 0
	
	for ok := true; ok; ok = (i < fat){
		if(n > 1){
			fat = fat*(n-1)
			fmt.Printf("fat: %d\n", fat)
			n--
		}
		i++
	}
	recursao(v2, n/4)
	recursao(v2, n/4)
	recursao(v2, n/4)
	recursao(v2, n/4)
	recursao(v2, n/4)
	recursao(v2, n/4)
	recursao(v2, n/4)
	recursao(v2, n/4)
}
func main(){
	v := make([]int, 12)
	recursao(v, 12)
}