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

var reader = bufio.NewReader(os.Stdin)

func CheckTaskInput() {

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
	case "task", "task help":
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task help" + " - Show help" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task list" + " - Show all tasks" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task add <Task>" + " - Add new task" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task remove <Task>" + " - Remove task" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task remove all" + " - Clear all tasks" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task edit <Task>" + " - Edit task" + Reset)
		//fmt.Println(Green+"task done"+Reset+" - Mark task as done")
		//fmt.Println(Green+"task undone"+Reset+" - Mark task as undone")
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task info <Task>" + Reset + " - Show task info" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task info all" + Reset + " - Show all tasks info" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task start <Task>" + Reset + " - Start task" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task start all" + Reset + " - Starts all tasks" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task stop all" + Reset + " - Stops all tasks" + Reset)
		fmt.Println(Green + "Task Help" + Reset + " » " + Blue + "task stop <Task>" + Reset + " - Stop task" + Reset)
	case "task list":
		fmt.Println("Listing all tasks...")
	case "task add":
		fmt.Println("Adding new task...")
	case "task remove":
		fmt.Println("Removing task...")
	case "task remove all":
		fmt.Println("Removing all tasks...")
	case "task edit":
		fmt.Println("Editing task...")
	case "task info":
		fmt.Println("Showing task info...")
	case "task info all":
		fmt.Println("Showing all tasks info...")
	case "task start":
		fmt.Println("Starting task...")
	case "task start all":
		fmt.Println("Starting all tasks...")
	case "task stop":
		fmt.Println("Stopping task...")
	case "task stop all":
		fmt.Println("Stopping all tasks...")

	/**case "your name":
	name, _ := getInput(">>", reader)
	name1 := name
	fmt.Println(name1)**/
	default:
		println(Yellow + opt + Red + " is not available. Type " + Yellow + "task help " + Red + "to see the commands.")
	}

}
