package main

import "fmt"

func teste(v []int, n int){
	if(n < 2){
		return
	}
	v1 := make([]int, n/2)
	v2 := make([]int, n/2)
	
	teste(v1, n/2)
	teste(v2, n/2)
}
func maior(v []int, n int){
	var inicio int = 0
	var fim int = n-1
	var meio int
	var maior int = v[n-1]
	
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
	v := make([]int, 10)
	v = []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	maior(v, 10)
	teste(v, 10)
}