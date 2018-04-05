package main

//	Imports
import (
	"fmt"
)	//	import

import "./Shunt"
import "./NFA"

func main() {
    Menu()
}   //  main()

//  Menu function handling user interaction
func Menu() {
    input := ""
    postfix := ""
	infix := "a.b.c*"
    s := "ab"

    Options()   //  display menu options

    //  Loop until sentinel gets inputted
    for input != "-1" {
        fmt.Print("\nYour Option (-1 to Exit, 0 for Menu Options): ")
        fmt.Scan(&input)    //  user input - choosing an option
        
        //  Switch handling different options
        switch input {
            case "-1":
                continue
            case "0":
                Options()   //  display options
            case "1":
                fmt.Print("Infix: ")
                fmt.Scan(&infix)    //  user input
            case "2":
                fmt.Print("String: ")
                fmt.Scan(&s)
            case "3":
                postfix = Shunt.ConvertToPostfix(infix)   //  convert to postfix

                fmt.Println("\nDoes the regular expression matches given string of text: ", NFA.PoMatch(postfix, s))   //  output
                fmt.Println("************************************************************")
            default:
                fmt.Println("Invalid Input, Please Select from Options 1, 2, 3 or -1 to Exit!") //  let the user know if wrong option was chosen
        }   //  switch
	}   //  for
}   //  Main()

//  List of menu options
func Options() {
    fmt.Println("\n************************************************************")
    fmt.Println("1. Parse Regular Expression in Infix Form.")
    fmt.Println("2. Parse String of Text.")
    fmt.Println("3. Check if the Regular Expression Matches Given String of Text.")
}   //  Options()