package main

import(
	"fmt"
	"strings"
	"bufio"
	"unicode"
	"os"
)

func main(){
	updatedmain()
}

func updatedmain() {
	fmt.Println("Enter a String: ")
	reader := bufio.NewReader(os.Stdin)
	input,_:= reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if len(input) == 0{
		fmt.Println("Not Found!")
		return
	}
	if len(input) < 3 {
		fmt.Println("Not Found!")
		return
	}

	ifoundStart, nfoundEnd, afound := false, false, false
	n := len(input)

	for i, char := range input {
		if i == 0 && unicode.ToLower(char) == 'i'{
			ifoundStart = true
		}

		if i == n-1 && unicode.ToLower(char) == 'n' {
			nfoundEnd = true
		}

		if unicode.ToLower(char) == 'a'{
			afound = true
		}

		if ifoundStart && nfoundEnd && afound {
			fmt.Println("Found!")
			return
		}
	}
	fmt.Println("Not Found!")
}
