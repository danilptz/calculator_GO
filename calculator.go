package main

import (
	"fmt"
	"strconv"
)

var rimNum = [11]string{"0", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
var num string
var a, b int

func main() {

	_, _ = fmt.Scan(&num)
	lst := [9]string{}
	count := 0
	var operIndex int
	var oper string
	var aStr string
	var bStr string

	//-------------------Записываем данные в массив-----------------------------------------

	for i := range num {
		lst[i] = num[i : i+1]
	}

	//-------------------Получаем строковые операнды и оператор-----------------------------

	for i := range lst {
		if lst[i] == "+" || lst[i] == "-" || lst[i] == "*" || lst[i] == "/" {
			count++
			if count == 1 {
				oper = lst[i]
				operIndex = i
			}
			for j := range lst {
				if j > operIndex {
					bStr += lst[j]
				}
			}

		}

		if count > 1 {
			fmt.Print("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			break
		}

		if count < 1 {
			aStr += lst[i]
		}
	}

	if count < 1 {
		fmt.Print("Вывод ошибки, так как строка не является математической операцией.")
	}

	//-----------------Обработка данных-------------------------------------------------------

	if rimOrArab(aStr) > 0 {
		a = rimToArab(aStr)
	} else {
		a, _ = strconv.Atoi(aStr)
	}

	if rimOrArab(bStr) > 0 {
		b = rimToArab(bStr)
	} else {
		b, _ = strconv.Atoi(bStr)
	}

	if rimOrArab(aStr) > 0 && rimOrArab(bStr) > 0 {
		fmt.Print(arabToRim(decision(oper)))
	} else if rimOrArab(aStr) < 1 && rimOrArab(bStr) < 1 {
		fmt.Print(decision(oper))
	} else if rimOrArab(aStr) < 1 && rimOrArab(bStr) > 0 || rimOrArab(aStr) > 0 && rimOrArab(bStr) < 1 {
		fmt.Print("Вывод ошибки, так как используются одновременно разные системы счисления.")
	} else {
		fmt.Print("Вывод ошибки, так как строка не является математической операцией.")
	}
}

// ------------------Проверка на римскую цифру-----------------------------------------------

func rimOrArab(num string) int {

	res := 0
	for i := range rimNum {
		if num == rimNum[i] {
			res++
		}
	}
	return res
}

// ------------------Перевод из римской цифры в арабскую--------------------------------------

func rimToArab(rim string) int {

	arab := 0
	for i := range rimNum {
		if rim == rimNum[i] {
			arab = i
		}
	}
	return arab
}

// ------------------Выполнение действия и вывод решения арабскими цифрами--------------------

func decision(operator string) int {
	var res int

	switch operator {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}
	return res
}

//------------------Перевед ответа в римскую цифру---------------------------------------------

func arabToRim(num int) string {
	var aInt, bInt int
	var res, aRim, bRim string

	if num > 10 && num < 100 {
		aInt = num / 10
		bInt = num % 10
	}

	switch aInt {
	case 1:
		aRim = "X"
	case 2:
		aRim = "XX"
	case 3:
		aRim = "XXX"
	case 4:
		aRim = "XL"
	case 5:
		aRim = "L"
	case 6:
		aRim = "LX"
	case 7:
		aRim = "LXX"
	case 8:
		aRim = "LXXX"
	case 9:
		aRim = "XC"
	}

	switch bInt {
	case 1:
		bRim = "I"
	case 2:
		bRim = "II"
	case 3:
		bRim = "III"
	case 4:
		bRim = "IV"
	case 5:
		bRim = "V"
	case 6:
		bRim = "VI"
	case 7:
		bRim = "VII"
	case 8:
		bRim = "VIII"
	case 9:
		bRim = "IX"
	}

	if num == 100 {
		res = "C"
	} else if 10 < num && num < 100 {
		res = aRim + bRim
	} else if num < 10 && num > 0 {
		res = bRim
	} else if num < 0 {
		res = "Вывод ошибки, так как в римской системе нет отрицательных чисел."
	}
	return res
}
