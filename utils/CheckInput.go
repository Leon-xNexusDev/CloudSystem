package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func CheckInput() {

	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Blue = "\033[34m"
	var White = "\033[97m"
	var Green = "\033[32m"
	var Yellow = "\033[33m"

	reader := bufio.NewReader(os.Stdin)
	hostName, _ := os.Hostname()
	userName, _ := user.Current()

	opt, _ := getInput(White+userName.Username+Reset+"@"+Red+hostName+Blue+" » ", reader)

	switch opt {
	case "cloud stop":
		fmt.Println("Stopping Cloud...")
		os.Exit(0)
	case "cloud help", "cloud":
		fmt.Println(Green + "Help" + Reset + " » " + "cloud stop")
	/*case "your name":
	name, _ := getInput(">>", reader)
	fmt.Println(name)*/
	default:
		println(Red + "This command is not available. Type " + Yellow + "cloud help " + Red + "to see the commands.")
	}

}
