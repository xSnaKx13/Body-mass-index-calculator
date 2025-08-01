package main

import (
	"fmt"
)

func result(result float64) string {
	var text string
	switch {
	case result < 16:
		text = fmt.Sprintln("У вас выраженный дефицит массы тела!")

	case result > 16 && result < 18.5:
		text = fmt.Sprintln("У вас недостаточная масса тела.")

	case result > 18.5 && result < 25:
		text = fmt.Sprintln("У вас нормальный вес.")

	case result > 25 && result < 30:
		text = fmt.Sprintln("У вас чрезмерная масса тела (предожирение).")

	case result > 30 && result < 35:
		text = fmt.Sprintln("У вас ожирение I-ой степени!")

	case result > 35 && result < 40:
		text = fmt.Sprintln("У вас ожирение II-ой степени!")

	case result > 40:
		text = fmt.Sprintln("У вас ожирение III-ей степени!")

	}
	return text
}
