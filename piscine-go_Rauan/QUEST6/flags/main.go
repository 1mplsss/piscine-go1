package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args[1:]
	ins := ""
	txt := ""
	ord, help, print := false, true, true
	for _, elem := range arr {
		if elem == "--order" || elem == "-o" {
			ord = true
		} else if elem == "--help" || elem == "-h" {
			fmt.Printf("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
			print = false
		} else if elem != "" && elem[:1] == "-" {
			if elem[:3] == "-i=" {
				ins += elem[3:]
			} else if elem[:9] == "--insert=" {
				ins += elem[9:]
			}
		} else {
			txt += elem
		}
		help = false
	}
	if help {
		fmt.Printf("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
		print = false
	}
	if !help && print {
		txt += ins
		if ord {
			arr := []rune(txt)
			for true {
				swapped := false
				for i := 0; i < len(arr)-1; i++ {
					if arr[i] > arr[i+1] {
						swapped = true
						arr[i], arr[i+1] = arr[i+1], arr[i]
					}
				}
				if !swapped {
					break
				}
			}
			for i := range arr {
				z01.PrintRune(arr[i])
			}
			z01.PrintRune('\n')
		} else {
			for i := range txt {
				z01.PrintRune(rune(txt[i]))
			}
			z01.PrintRune('\n')
		}
	}
}
