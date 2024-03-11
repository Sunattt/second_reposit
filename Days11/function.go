package Days11

import "fmt"

func TransStr(t string) []uint8 {
	slice := []uint8{}
	for i := 0; i < len(t); i++ {
		fmt.Printf("%v ", t[i])
		slice = append(slice, t[i])
 	}
	fmt.Println()
	return slice
}

func TransByte(b []uint8) []string{
	var array []string
	for _, l := range b {
		f := fmt.Sprintf("%c", l)
		fmt.Printf("%c", l)
		array = append(array, f)
	}
	fmt.Println()
	return array
}

func Caezar(str string) {
	for i := 0; i < len(str); i++ {
		letter := str[i] + 3
		if letter > 122 {
			letter = letter - 122 + 96
		} else if letter == 32 {
			letter += 3
		}
		fmt.Printf("%c", letter)
		fmt.Print(" - ")
		fmt.Println(letter)
	}
}

func ReplaceSim(text string) {
	/* это один из способов решение 
	text1 := strings.Replace(text, "+", " ", -1) 
	*/
	for _ , el := range text {
		if el == 43 {
			fmt.Print(" ")
		} else {
		fmt.Printf("%c", el )
	}
}
}
