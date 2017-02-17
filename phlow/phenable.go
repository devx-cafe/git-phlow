package phlow

import (
	"bufio"
	"fmt"
	"os"
)

//Enable ...
func Enable() {

	//var username = ""
	//var password = ""

	//executor := subprocess.SimpleExec
	//config := gitwrapper.NewConfigurator()

	fmt.Fprintf(os.Stdout, "enter credentials for %s \n", "github")
	//Read user input username
	ReadInput("username: ")
	//Read user input password
	ReadInput("password: ")

}

//ReadInput ...
//Reads input from user
func ReadInput(messageToUser string) string {
	scanner := bufio.NewReader(os.Stdin)
	text, _ := scanner.ReadString('\n')

	return text
}
