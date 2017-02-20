package phlow

import (
	"bufio"
	"fmt"
	"os"
)

/*
   Reauthorize should also be possible
*/

//Enable ...
func Enable() {
	//Run status before check

	token, tErr := Config("token", "", true)
	user, uErr := Config("user", "", true)

	if (tErr == nil && token != "") || (uErr == nil && user != "") {
		fmt.Println("you are already signed in")
		return
	}

	fmt.Fprintf(os.Stdout, "enter credentials for %s \n", "github")
	//Read user input username
	username := ReadInput("username: ")
	//Read user input password
	password := ReadInput("password: ")

	if token, err := Authorize(username, password, authURL); err != nil {
		Config("token", token, false)
		Config("user", username, false)
		fmt.Println("Success fully authorized: 'git do' is now enabled  ")
	}
}

//ReadInput ...
//Reads input from user
func ReadInput(messageToUser string) string {
	scanner := bufio.NewReader(os.Stdin)
	text, _ := scanner.ReadString('\n')

	return text
}
