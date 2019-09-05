package main

func recursao(v []int, n int){
	if(n < 4){
		return
	}
	v1 := make([]int, n/4)
	v2 := make([]int, n/4)
	var count int = 0
	recursao(v1, n/4)
	recursao(v2, n/4)
	recursao(v1, n/4)
	recursao(v2, n/4)
	recursao(v1, n/4)
	recursao(v2, n/4)
	recursao(v1, n/4)
	recursao(v2, n/4)
	for i := 0; i < n; i++{
		count++
	}	
	recursao(v1, n/4)
	recursao(v2, n/4)
	recursao(v1, n/4)
	recursao(v2, n/4)
	recursao(v1, n/4)
	recursao(v2, n/4)
	recursao(v1, n/4)
	recursao(v2, n/4)
}
func main(){
	v := make([]int, 10)
	v = []int{5, 8, 6, 7, 3, 1, 0, 2, 9, 4}
	recursao(v, 10)
}