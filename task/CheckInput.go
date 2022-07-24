package task

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

func CheckTaskInput() {

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
	case "task", "task help":
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task help" + " - Show help" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task list" + " - Show all tasks" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task add <Task>" + " - Add new task" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task remove <Task>" + " - Remove task" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task remove all" + " - Clear all tasks" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task edit <Task>" + " - Edit task" + Reset)
		//fmt.Println(Green+"task done"+Reset+" - Mark task as done")
		//fmt.Println(Green+"task undone"+Reset+" - Mark task as undone")
		fmt.Println(Green + "task info <Task>" + Reset + " - Show task info")
		fmt.Println(Green + "task info all" + Reset + " - Show all tasks info")
		fmt.Println(Green + "task start <Task>" + Reset + " - Start task")
		fmt.Println(Green + "task start all" + Reset + " - Starts all tasks")
		fmt.Println(Green + "task stop all" + Reset + " - Stops all tasks")
		fmt.Println(Green + "task stop <Task>" + Reset + " - Stop task")
	/*case "your name":
	name, _ := getInput(">>", reader)
	fmt.Println(name)*/
	default:
		println(Red + "This command is not available. Type " + Yellow + "task help " + Red + "to see the commands.")
	}

}
