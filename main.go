package main

import (
	"bufio"
 	"fmt"
 	"os"
	"strings"
)

func calculator (str string) string {
	var rez string
	var poz int

	strRunes := []rune(str)
	poz = 0

	for i := 0; i < len(strRunes); i++ {
		if strRunes[i] == '+' || strRunes[i] == '-' || strRunes[i] == '/' || strRunes[i] == '*'{
			poz = i
		} 
	}

		if 	(strRunes[0] != '"' ||
			strRunes[poz-2] != '"' ||
			strRunes[poz-1] != ' ' ||
			strRunes[poz+1] != ' ' ||
			poz < 4) {
				rez = "Неверный формат 2"
		} else {

			lenstr2 := (len(strRunes)-3) - (poz+1)

			if strRunes[poz] == '+' || strRunes[poz] == '-' {
				if strRunes[poz+2] != '"' || strRunes[len(strRunes)-3] != '"' || lenstr2 < 3 {
					rez = "Неверный формат 3"
				}
				if strRunes[poz] == '+' {
					str1 := str[1:(poz-2)]
					str2 := str[poz+3:len(str)-3]
					rez = "\"" + str1 + str2 + "\""	
				} else {
					str1 := str[1:(poz-2)]
					str2 := str[poz+3:len(str)-3]
					index := strings.Index(str1, str2)
					if index != -1 {
						rez = "\"" + str1[:index] + str1[index+len(str2):] + "\""
					} else {
						rez = "\"" + str1 + "\""	
					}

				}
			} else {
				if (lenstr2 > 2) || (lenstr2 < 1) {
					rez = "Неверный формат 4"
				} else {
					if lenstr2 == 1 {
						switch strRunes[len(strRunes)-3] {
							case '0':
								rez = "Неверный формат 5"
							case '1':
								
							case '2':
								
							case '3':
								
							case '4':
								
							case '5':
								
							case '6':
								
							case '7':
								
							case '8':
								
							case '9':
								
						}
					}
					if lenstr2 == 2 {
						if strRunes[len(strRunes)-4] == '1' && strRunes[len(strRunes)-3] == '0' {
							
						} else {
							rez = "Неверный формат 6"
						}
					}
					
				}
			}
		}


	

	return rez
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var str string

	fmt.Println("Введи выражение")

	for {
		str, _ = reader.ReadString('\n')
		rez := calculator(str)
		fmt.Println(rez)

	}			
		
	

	



}


	
	
