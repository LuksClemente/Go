package main

import "fmt"

func calcular(a int) (quadradro int, cubo int){
	quadradro = a*a
	cubo = a*a*a
	
	return
}
func main(){
	fmt.Println(calcular(2))
}