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

var reader = bufio.NewReader(os.Stdin)

func CheckInput() {

	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Blue = "\033[34m"
	var White = "\033[97m"
	var Green = "\033[32m"
	var Yellow = "\033[33m"

	hostName, _ := os.Hostname()
	userName, _ := user.Current()

	opt, _ := getInput(White+userName.Username+Reset+"@"+Red+hostName+Blue+" » ", reader)

	switch opt {
	case "cloud help", "cloud":
		fmt.Println(Green + "Help" + Reset + " » " + Blue + "cloud help (shows this help)" + Reset)
		fmt.Println(Green + "Help" + Reset + " » " + Blue + "cloud version (shows the version of the CloudSystem)" + Reset)
		fmt.Println(Green + "Help" + Reset + " » " + Blue + "cloud info (shows the info of the CloudSystem)" + Reset)
		fmt.Println(Green + "Help" + Reset + " » " + Blue + "cloud stop (stops the CloudSystem)" + Reset)
		fmt.Println(Green + "Help" + Reset + " » " + Blue + "task help (shows the help of the task command)" + Reset)
	case "cloud stop":
		fmt.Println("Stopping Cloud...")
		os.Exit(0)
	case "cloud version":
		fmt.Println("CloudSystem v1.0 alpha")
	case "cloud info":
		fmt.Println("Version v1.0 alpha")
		fmt.Println("Author: Leon-xNexusDev & Gamer_Bjoern")
		fmt.Println("Discord: https://discord.gg/QWQWQWQ") //TODO Add discord link
		fmt.Println("License: GNU GPLv3")
		fmt.Println("Copyright: 2022")
		fmt.Println("All rights reserved.")
	/*case "your name":
	name, _ := getInput(">>", reader)
	fmt.Println(name)*/
	default:
		println(Red + "This command is not available. Type " + Yellow + "cloud help " + Red + "to see the commands.")
	}

}
