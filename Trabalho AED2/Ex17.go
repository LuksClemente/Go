package main

import "fmt"
import "math"

func recursao(v []int, n int){
	if(n < 3){
		return
	}
	v1 := make([]int, n/3) 
	v2 := make([]int, n/3)
	v3 := make([]int, n/3)
	var cont int = 0
	for i := 0; i < n; i++{
		for j := 0; i < n; j++{
			v[i] = cont
			v[j] = cont
			cont++
			for k := 0; float64(k) < (math.Log(float64(n))); k++{
				v[j] = v[j] + 5
			}
		}
	}
	recursao(v1, n/3)
	recursao(v2, n/3)
	recursao(v3, n/3)
}
func main(){
	v := make([]int, 15)
	recursao(v, 15)
	for i := 0; i < 15; i++{
		fmt.Printf("%d -> ", v[i])
	}
}