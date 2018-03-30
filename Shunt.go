package main

//	Imports
import (
	"fmt"
)	//	import

//	Convert infix to pofix
func intopost(infix string) string {
	//	Map special characters (e.g. *) stored as integers
	specials := map[rune]int{'*':10, '.':9, '|': 8}	//	keeping track of precedence of special characters
	
	//	Array of runes of empty
	pofix := []rune{}	//	runes - character as displayed on the screen
	s: = []rune{}	//	stack of operators from infix of regular expression
	
	return string(pofix)
}	//	intopost()

func main() {
	//	Answer: ab.c*.
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Postfix: ", intopost("a.b.c*"))

	//	Answer: abd|.*
	fmt.Println("Infix: ", "(a.(b|d))*")
	fmt.Println("Postfix: ", intopost("(a.(b|d))*"))

	//	Answer: abd|.c*.
	fmt.Println("Infix: ", "(a.(b|d)).c*")
	fmt.Println("Postfix: ", intopost("(a.(b|d)).c*"))

	//	Answer: abb.+.c.
	fmt.Println("Infix: ", "a.(b.b)+.c")
	fmt.Println("Postfix: ", intopost("a.(b.b)+.c"))

}	//	main()