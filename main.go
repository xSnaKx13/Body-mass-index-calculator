package main

import "fmt"

func main() {
	for {
		var weight, height float64

		err := userInput(&weight, &height)
		if err != nil {
			fmt.Print(err)
		}
		IMTResult := calculation(weight, height)
		textResult := result(IMTResult)
		fmt.Printf("Ваш ИМТ: %.2f: %s", IMTResult, textResult)

		for {
			fmt.Print("1 - расчитать повторно\n2 - выход\n")
			var input int
			fmt.Scan(&input)

			switch input {
			case 1:
				continue
			case 2:
				return
			default:
				fmt.Println("Неверный ввод, введите 1 или 2")
			}
		}
	}
}
