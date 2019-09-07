package main

import "fmt"

var qtd int = 0

func fibonacci(n int) int{
    if n <= 1{
		return n
	}else{
		qtd += 2
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
func main(){
	var n int
	fmt.Printf("Digite um numero qualquer: ")
	fmt.Scanf("%d", &n)
	fmt.Printf("%d° numero da sequencia: %d\n", n, fibonacci(n))
	fmt.Printf("Quantidade de chamadas recursivas: %d", qtd)
}