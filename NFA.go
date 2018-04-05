package main

//	Imports
import (
	"fmt"
)	//	import

//  Max arrows from a state is two
type state struct {
    symbol rune //  instead of labeling the arrows

    //  Pointers to other states
    edgeOne *state    //  arrow one
    edgeTwo *state    //  arrow two
}   //  state

//  Thompson's construction - for each fragment there is one initial state and one accept state
//  Keep track of initial state and accept state of fragment of nfa
type nfa struct {
    initial *state
    accept *state
}   //  nfa

//  Postfix regular expression to non deterministic finite automaton
func ConvertToNFA(pofix string) *nfa {
    //  Thompson's construction work on a stack to keep a track of fragments of nfa's on a stack
    nfaStack := []*nfa{}

    //  Loop through postfix regular expression, rune at a time
    for _, r := range pofix {
        switch r {
            case '.':
                //  Pop two thing off nfa stack - two pointers to nfa fragments
                fragmentTwo := nfaStack[len(nfaStack) - 1]  //  last stack
                nfaStack = nfaStack[:len(nfaStack) - 1] //  //  get rid off last thing thats on the stack

                fragmentOne := nfaStack[len(nfaStack) - 1]  //  last stack
                nfaStack = nfaStack[:len(nfaStack) - 1] //  //  get rid off last thing thats on the stack
                
                //  Join two fragments together and then push concatenated fragment back to nfa stack
                fragmentOne.accept.edgeOne = fragmentTwo.initial    //  first edge of the accept state point to edge two of initial state
                
                nfaStack = append(nfaStack, &nfa{initial: fragmentOne.initial, accept: fragmentTwo.accept}) //  append instance of address of nfa structs
            case '|':
                //  Pop two thing off nfa stack - two pointers to nfa fragments
                fragmentTwo := nfaStack[len(nfaStack) - 1]  //  last stack
                nfaStack = nfaStack[:len(nfaStack) - 1] //  //  get rid off last thing thats on the stack

                fragmentOne := nfaStack[len(nfaStack) - 1]  //  last stack
                nfaStack = nfaStack[:len(nfaStack) - 1] //  //  get rid off last thing thats on the stack
                
                //  Create new initial state and accept state, then join the states to the fragments
                accept := state{}
                initial := state{edgeOne: fragmentOne.initial, edgeTwo: fragmentTwo.initial}

                //  Joins all the arrows in the correct way - pass new points
                fragmentOne.accept.edgeOne = &accept
                fragmentTwo.accept.edgeOne = &accept
                
                nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})   //   push new stack
            case '*':
                //  Pop one thing off nfa stack
                fragment := nfaStack[len(nfaStack) - 1]  //  last stack
                nfaStack = nfaStack[:len(nfaStack) - 1] //  //  get rid off last thing thats on the stack

                //  Create new initial state and accept state, then join the states to the fragments
                accept := state{}
                initial := state{edgeOne: fragment.initial ,edgeTwo: &accept}

                //  Joins all the arrows in the correct way - pass new points
                fragment.accept.edgeOne = fragment.initial
                fragment.accept.edgeTwo = &accept

                nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})   //   push new stack
            default:
                accept := state{}
                initial := state{symbol: r, edgeOne: &accept}

                nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})   //   push new stack
        }   //  switch
    }   //  for

    return nfaStack[0]
}   //  poregtonfa()

//  Main
func main() {
    nfa := ConvertToNFA("ab.c*|")
    fmt.Println(nfa)
}   //  main()