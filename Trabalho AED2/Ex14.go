package main

func recursao(v []int, n int){
	if(n < 2){
		return
	}
	v1 := make([]int, n/2) 
	v2 := make([]int, n/2)
	recursao(v1, n/2)
	recursao(v2, n/2)
	var k int = 2
	
	for i := 0; i < k*n; i++{
		if(n/2 > i){
			v1[i] = 0
			v2[i] = 0
		}
	}
	recursao(v1, n/2)
	recursao(v2, n/2)
}
func main(){
	v := make([]int, 12)
	v = []int{5, 8, 11, 2, 25, 17, 10, 7, 9, 28, 32, 12}
	recursao(v, 3)
}