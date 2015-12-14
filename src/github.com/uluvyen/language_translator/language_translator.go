package main

import "fmt"

func main() {
	var characters = getallenglishcharacters()


	// waits for the user to request to leave.
	waitforexit()
}

func waitforexit() {
	
	exit_string := "exit"
	
	fmt.Print("Please enter exit to leave the program: ")
	var input string
	fmt.Scanln(&input)

	for (input != exit_string) {
		fmt.Print("INVALID VALUE... Please enter exit to leave the program: ")
		fmt.Scanln(&input)
	}
}

func getallenglishcharacters() [27]string {
	var characters [27] string
	var characterstring = "abcdefghijklmnopqrstuvwxyz"
	
	for i, char := range characterstring {
		// save c in spot i
		characters[i] = string(char)
	} 
	

	return characters
}