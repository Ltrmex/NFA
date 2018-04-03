package main

//	Imports
import (
	"fmt"
)	//	import

//  Max arrows from a state is two
type state struct {
    symbol rune //  instead of labeling the arrows

    //  Pointers to other states
    edge1 *state    //  arrow one
    edge2 *state    //  arrow two
}   //  state

//  Thompson's construction - for each fragment there is one initial state and one accept state
//  Keep track of initial state and accept state of fragment of nfa
type nfa struct {
    initial *state
    accept *state
}   //  nfa

//  Postfix regular expression to non deterministic finite automaton
func ConvertToNFA(pofix string) *nfa {

}   //  poregtonfa()

//  Main
func main() {
    nfa := ConvertToNFA("ab.c*|")
    fmt.Println(nfa)
}   //  main()