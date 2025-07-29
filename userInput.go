package main

import (
	"fmt"
)
func userInput(userHeight, userWeight float64, err error)(float64 , float64, error){
	fmt.Print("Введите ваш вес (кг): ")
	_, err = fmt.Scan(&userWeight)
	if err!=nil {
		fmt.Println("Введены невалидные данные, укажите вес корректно!")
		return 0,0,err
	}
	fmt.Print("Введите ваш рост (см): ")
	_, err = fmt.Scan(&userHeight)
	if err!=nil {
		fmt.Println("Введены невалидные данные, укажите рост правильно!")
		return 0,0,err
	}
	return userWeight, userHeight, nil
}
