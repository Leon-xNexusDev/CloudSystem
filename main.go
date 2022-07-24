package main

import (
	"github.com/Leon-xNexusDev/CloudSystem/task"
	"github.com/Leon-xNexusDev/CloudSystem/utils"
)

func main() {

	utils.SendASCI()
	for {
		utils.CheckInput()
		task.CheckTaskInput()
	}

}
