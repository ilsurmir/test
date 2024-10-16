package main

import (
	"bufio"
	"os"
 	"fmt"
 	"strings"
)

func getThePartsAndTypeOfOperation (str string) (str1, operation, str2 string) {

	if strings.Contains(str, "\" + \"") {
		operation = "+"
		index := strings.Index(str, "\" + \"")
		str1 = str[:index+1]
		str2 = str[index+len([]rune("\" + \""))-1:]
	} else {
		if strings.Contains(str, "\" - \"") {
			operation = "-"
			index := strings.Index(str, "\" - \"")
			str1 = str[:index+1]
			str2 = str[index+len([]rune("\" - \""))-1:]
		} else {
			if strings.Contains(str, "\" * ") {
				operation = "*"
				index := strings.Index(str, "\" * ")
				str1 = str[:index+1]
				str2 = str[index+len([]rune("\" * ")):]
			} else {
				if strings.Contains(str, "\" / ") {
					operation = "/"
					index := strings.Index(str, "\" / ")
					str1 = str[:index+1]
					str2 = str[index+len([]rune("\" / ")):]
				} else {
					panic("Упадок программы")
				}
			}
		}
	}

	

	return str1, operation, str2

}

func checkingForErrors (str1, operation, str2 string) (bool) {
	
	if 	len([]rune(str1)) < 3 || 
		(strings.Index(str1, "\"")) != 0 || 
		len([]rune(str1)) > 10 {
		return false 
	}

	if (operation == "*" || operation == "/") {
		
		if 	str2 != "1" &&
			str2 != "2" &&
			str2 != "2" &&
			str2 != "3" &&
			str2 != "4" &&
			str2 != "5" &&
			str2 != "6" &&
			str2 != "7" &&
			str2 != "8" &&
			str2 != "9" &&
			str2 != "10" {
				return false 
		}
	} else {
		if 	len([]rune(str2)) < 3 || 
			str2[strings.LastIndex(str2, "\"")] != '"' ||
			len([]rune(str2)) > 10 {
			
				return false 
		}
	}

	return true 
}

func summation (str1, str2 string) (result string) {
	
	result = str1[:len(str1)-1] + str2[1:]

	return result
}

func subtraction (str11, str22 string) (result string) {

	str1:= str11[1:len(str11)-1]
	str2:= str22[1:len(str22)-1]

	index := strings.Index(str1, str2)

	if index != -1 {
		result = "\"" + str1[:index] + str1[index+len(str2):] + "\""
	} else {
		result = str11	
	}

	return result
}

func str2ToInt (str2 string) (str2int int) {
	switch str2 {
		case "1":
			str2int = 1
		case "2":
			str2int = 2
		case "3":
			str2int = 3
		case "4":
			str2int = 4
		case "5":
			str2int = 5
		case "6":
			str2int = 6
		case "7":
			str2int = 7
		case "8":
			str2int = 8
		case "9":
			str2int = 9
		case "10":
			str2int = 10
	}

	return str2int
}

func multiplication (str11, str22 string) (result string) {

	str1:= str11[1:len([]rune(str11))-1]
	str2 := str2ToInt(str22)

	for i := 0; i < str2; i++ {

		result += str1
	}
	result = "\"" + result + "\""

	return result
}

func division (str11, str22 string) (result string) {

	str1:= str11[1:len(str11)-1]
	str2 := str2ToInt(str22)

	result = str1[:(len(str1)) / str2]
	result = "\"" + result + "\""

	return result
}


func main() {

	var result string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение")
	
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	str1, operation, str2 := getThePartsAndTypeOfOperation(str)

	allowance := checkingForErrors(str1, operation, str2)
	if allowance == false {
		panic("Упадок программы")
	}

	switch operation {
		case "+":
			result = summation (str1, str2)
		case "-":
			result = subtraction (str1, str2)
		case "*":
			result = multiplication (str1, str2)
		case "/":
			result = division (str1, str2)
	}

	if (len(result) <= 42) {
		fmt.Println(result)
	} else {
		fmt.Println(result[:41] + "...\"")
	}
}