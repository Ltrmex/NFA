/* 
Code Reference: https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
Note: Code was implemented through the use of the above tutorial
*/

package Shunt

//	Convert infix to pofix
func ConvertToPostfix(infix string) string {
	//	Map special characters (e.g. *) stored as integers
	specials := map[rune]int{'*':10, '+':9, '?':8, '.':6, '|':5}	//	keeping track of precedence of special characters

	//	Array of runes of empty
	pofix := []rune{}	//	runes - character as displayed on the screen
	s := []rune{}	//	stack of operators from infix of regular expression
	
	//	Loop over infix
	for _, r := range infix {	//	converts each part of the string into array of runes
		switch {
			case r == '(':
				s = append(s, r)	//	append to the end of the stack
			case r == ')':
				for s[len(s) - 1] != '(' {	//	last element of s
					pofix = append(pofix, s[len(s) - 1])	//	append what is at the top of the stack
					
					//	Get rid of the last element on the stack
					s = s[:len(s) - 1]	//	keep everything up to (excluding) last element
				}	//	for
				
				//	Get rid of open bracket
				s = s[:len(s) - 1]
			case specials[r] > 0:	//	check if current character is a special character, if try to access element that is not in the specials, return 0
				//	While still has elements && precedence is less or equal to precedence of the element at the top of the stack
				for len(s) > 0 && specials[r] <= specials[s[len(s) - 1]] {
					//	Take elements off the top of the stack and put into pofix
					pofix = append(pofix, s[len(s) - 1])
					s = s[:len(s) - 1]	
				}	//	for

				//	When element at the top of the stack has less precedence from current character
				s = append(s, r)	//	append current character
			default:
				pofix = append(pofix, r)	//	take rune r and add it to the end
		}	//	switch
	}	//	for

	//	If anything left on the top of the stack
	for len(s) > 0 {
		//	Take elements off the top of the stack and put into pofix
		pofix = append(pofix, s[len(s) - 1])
		s = s[:len(s) - 1]	
	}	//	for
	return string(pofix)
}	//	intopost()
