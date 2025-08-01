package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func userInput(weight *float64, height *float64) (err error) {

	for {
		fmt.Print("Введите ваш вес: ")
		weightStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		weightStr = strings.TrimSpace(weightStr)
		*weight, err = strconv.ParseFloat(weightStr, 64)
		if err != nil {
			fmt.Println("Некорректный вес, введите число")
			continue
		}

		fmt.Print("Введите ваш рост: ")
		heightStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		heightStr = strings.TrimSpace(heightStr)
		*height, err = strconv.ParseFloat(heightStr, 64)
		if err != nil {
			fmt.Println("Некорректный рост, введите число")
			continue
		}
		return nil
	}
}
