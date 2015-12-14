package main

import "fmt"

func main() {
	var characters = getallenglishcharacters()

	fmt.Println("Please input the string you wish to convert: ")
	var input string
	fmt.Scanln(&input)

	var indexlist []int
	for _, c := range input {
		var cString = string(c)
//		fmt.Println("Checking ", cString)
		var index = getindex(characters, cString)
//		fmt.Println("Character's index is ", index)
		indexlist = append(indexlist, index)
	}

	// for each element in the array, determine the image
		

	// waits for the user to request to leave.
	waitforexit()
}

func getindex(characterlist [27]string, input string) int {
	
	var index int

	for i, c := range characterlist {
		if c == input {
			index = i
		}
	}

	return int(index)
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