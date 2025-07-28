package main

import (
"fmt"
"math"
)

func main(){
	var weight float64 
	var height float64 
	fmt.Print("Введите ваш вес: ")
	fmt.Scan(&weight)
	fmt.Print("Введите ваш роств см: ")
	fmt.Scan(&height)
					  IMT := weight / math.Pow(height/100, 2)

					    fmt.Printf("Ваш ИМТ: %.2f", IMT)
						}
