# Graph Theory Project - Non-Deterministic Fnite Automaton

### NOTE: Make sure to check out the Wiki tab for more theory information about this project and methods connected to it

### Requirement
> You must write a program in the Go programming language [2] that can build a non-deterministic ﬁnite automaton (NFA) from a regular expression, and can use the NFA to check if the regular expression matches any given string of text. You must write the program from scratch and cannot use the regexp package from the Go standard library nor any other external library. 
- > Graph Theory Project Instructions

### Required To Run
* Golang
    - https://golang.org/doc/install
* Github
    - https://help.github.com/articles/set-up-git/

### How To Run
1. Open command prompt
    - By going into a search bar type "cmd"
    - Navigate to where you want to store the project
2. Clone repository by typing:
    - git clone https://github.com/Ltrmex/NFA.git
3. Navigate to the cloned folder
    - cd NFA
4. Run the project by typing following command:
    - go run Main.go
5. Follow on screen instructions

### Example
* By selecting option 1 input infix regular expression
    - Example: a.b.c*
* By selecting option 2 input string to compare against it
    - Example: abc
* By selecting option 3 you should get the answer false or true depending on if the string was matched along with postfix form of the regular expression
    - Note: The answer to the above example should be
        - Postfix: ab.c*
        - Matched: true
    - Note: Try selecting option 2 but input "ad" this time
        - Notice that the answer for Matched should be false this time
        
### References
* https://swtch.com/~rsc/regexp/regexp1.html
* Sets of videos for this project: https://web.microsoftstream.com/channel/f9970e30-b336-4145-8af3-a2bbe2938f5e
