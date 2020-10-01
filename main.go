package main

import (
	"fmt"
)

func main() {
	var dia int
	var mes int
	fmt.Scanln(&dia)
	fmt.Scanln(&mes)
	if dia < 1 {
		fmt.Println("dia no valido")
	}
	if mes%2 == 0 {
		if dia > 30 {
			fmt.Println("dia no valido")
		}
	} else {
		if dia > 31 {
			fmt.Println("dia no valido")
		}
	}
	switch mes {
	case 1:
		if dia <= 20 {
			fmt.Println("capricornio")
		} else {
			fmt.Println("acuario")
		}
	case 2:
		if dia <= 19 {
			fmt.Println("acuario")
		} else {
			fmt.Println("piscis")
		}
	case 3:
		if dia <= 20 {
			fmt.Println("piscis")
		} else {
			fmt.Println("aries")
		}
	case 4:
		if dia <= 20 {
			fmt.Println("aries")
		} else {
			fmt.Println("tauro")
		}
	case 5:
		if dia <= 21 {
			fmt.Println("tauro")
		} else {
			fmt.Println("geminis")
		}
	case 6:
		if dia <= 21 {
			fmt.Println("geminis")
		} else {
			fmt.Println("cancer")
		}
	case 7:
		if dia <= 23 {
			fmt.Println("cancer")
		} else {
			fmt.Println("leo")
		}
	case 8:
		if dia <= 23 {
			fmt.Println("leo")
		} else {
			fmt.Println("virgo")
		}
	case 9:
		if dia <= 23 {
			fmt.Println("virgo")
		} else {
			fmt.Println("libra")
		}
	case 10:
		if dia <= 23 {
			fmt.Println("libra")
		} else {
			fmt.Println("escorpio")
		}
	case 11:
		if dia <= 22 {
			fmt.Println("escorpio")
		} else {
			fmt.Println("sagitario")
		}
	case 12:
		if dia <= 21 {
			fmt.Println("sagitario")
		} else {
			fmt.Println("capricornio")
		}
	default:
		fmt.Println("mes no valido")
	}
}
