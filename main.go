package main

import (
	"fmt"

	"github.com/KayoRonald/exemplo-go/calc"
)

func main() {
  var nome string
	var num1 int
  var num2 int

  fmt.Print("Seu nome: ")
	fmt.Scan(&nome)

  fmt.Print("Digite o primeiro valor: ")
	fmt.Scan(&num1)
  
  fmt.Print("Digite o segundo valor: ")
	fmt.Scan(&num2)
  
  result := calc.Muilt(num1, num2)
	fmt.Printf("Olá, %v. Sua multiplicação entre %v x %v = %v \n", nome,num1, num2, result)
}

// func calc(num1 int, num2 int) int{
//   return num1 * num2
// }
