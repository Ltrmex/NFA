/* 
Code Reference: https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
                https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
                https://swtch.com/~rsc/regexp/regexp1.html
Note: Code was implemented through the use of the above tutorial
*/

package NFA
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
            case '?':
                //  Pop one thing off nfa stack
                fragment := nfaStack[len(nfaStack) - 1] //  last stack
                nfaStack = nfaStack[:len(nfaStack) - 1] //  get rid off last thing thats on the stack

                //  Create new initial state, then join the state to the fragments
                initial := state{edgeOne: fragment.initial, edgeTwo: fragment.accept}

                nfaStack = append(nfaStack, &nfa{initial: &initial, accept: fragment.accept})   //   push new stack
            case '+':
                //  Pop one thing off nfa stack
                fragment := nfaStack[len(nfaStack) - 1] //  last stack
                nfaStack = nfaStack[:len(nfaStack) - 1] //  get rid off last thing thats on the stack

                //  Create new initial state and accept state, then join the states to the fragments
                accept := state{}
                initial := state{edgeOne: fragment.initial, edgeTwo: &accept}
                
                //  Joins all the arrows in the correct way - pass new points
                fragment.accept.edgeOne = &initial

                nfaStack = append(nfaStack, &nfa{initial: fragment.initial, accept: &accept})   //   push new stack
            default:
                accept := state{}
                initial := state{symbol: r, edgeOne: &accept}

                nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})   //   push new stack
        }   //  switch
    }   //  for

    if len(nfaStack) != 1 {
        fmt.Println("Uh oh: ", len(nfaStack), nfaStack)
    }   //  if

    return nfaStack[0]
}   //  poregtonfa()

//  Helper function - find all states
func AddState(l []*state, s *state, a *state) []*state {
    l = append(l, s)    //  append s state to l
    
    //  Check s if there's e arrows going from it
    if s != a && s.symbol == 0 {
        l = AddState(l, s.edgeOne, a)   //  follow the first edge

        if s.edgeTwo != nil {   //  check if there's a second edge
            l = AddState(l, s.edgeTwo, a)   //  follow the second edge
        }   //  inner if
    }   //  if

    return l
}   //  AddState()

//  Check if regular expression matched a string
func PoMatch(po string, s string) bool {
    isMatch := false    //  set to false
    poNFA := ConvertToNFA(po)

    //  Track of states
    current := []*state{}   //  track of current states
    next := []*state{}  //  where current state can get to

    current = AddState(current[:], poNFA.initial, poNFA.accept)

    //  Generate next from current by looping each rune and check current state
    for _, r := range s {
        for _, c := range current { //  loop through each state
            if c.symbol == r {  //  check if set to r
                next = AddState(next[:], c.edgeOne, poNFA.accept)
            }   //  if
        }   //  inner for

        current = next  //  current is the next set of states
        next = []*state{}   //  reset value of next to empty
    }   //  for

    //  Check if any of states are the accept state
    for _, c := range current {
        if c == poNFA.accept {
            isMatch = true
            break;
        }

    }   //  for
    return isMatch
}   //  pomatch()

/*
//  Main
func main() {
    nfa := ConvertToNFA("ab.c*|")
    fmt.Println(nfa)

    fmt.Println(PoMatch("ab.c*|", "ab"))
}   //  main()
*/