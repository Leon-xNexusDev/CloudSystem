package main

import (
	"github.com/Leon-xNexusDev/CloudSystem/task"
	"github.com/Leon-xNexusDev/CloudSystem/utils"
)

func main() {

	utils.SendASSCI()
	for {
		utils.CheckInput()
		task.CheckTaskInput()
	}

}
