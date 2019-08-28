package main

import "fmt"

func main(){
	var a int = 46
	var b float64 = 6.4
	
	var c int = int(b)
	var d float64 = float64(a)
	
	fmt.Println("O valor de c é", c)
	fmt.Println("o valor de d é", d)
}
	