package main

import "fmt"

func recursao(v []int, n int){
	if(n < 4){
		return
	}
	v1 := make([]int, n/4) 
	v2 := make([]int, n/4)
	v3 := make([]int, n/4)
	
	recursao(v1, n/4)
	recursao(v2, n/4)
	recursao(v3, n/4)
}
func maior(v []int, n int){
	var inicio int = 0
	var fim int = n - 1
	var meio int
	var maior = v[n-1]
	
	for ok := true; ok; ok = (inicio <= fim){
		meio = (inicio + fim)/2
		if(maior == v[meio]){
			maior = v[meio]
			break
		}
		if(maior < v[meio]){
			fim = meio - 1
		}else{
			inicio = meio + 1
		}
	}
	fmt.Printf("Maior: %d\n", maior)
	for i := 0; i < n; i++{
		fmt.Printf("%d -> ", v[i])
	}
	fmt.Printf("\n")
}
func main(){
	v := make([]int, 16)
	v = []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 30, 34, 35, 40, 41, 45}
	maior(v, 16)
	recursao(v, 16)
}