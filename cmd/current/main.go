package main

import (
	"strings"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"github.com/jubnzv/go-taskwarrior"
	"github.com/zacharyfrederick/current-next"
	"flag"
)

func main() {
	setPtr := flag.Bool("s", false, "whether to set the current task to the supplied task_id")
	flag.Parse()

	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Your home directory is not set")
		os.Exit(3)
	}
	
	data_path := filepath.Join(homedir, ".current-next", "current.data")
	if _, err := os.Stat(data_path); os.IsNotExist(err) {
		fmt.Println("The current task has not been set") 
		os.Exit(3)
	}

	data, err := os.ReadFile(data_path)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	temp := strings.Replace(string(data), "\n", "", -1)
	if temp == "" {
		fmt.Println("No current task set")
		os.Exit(0)
	}
	task_id, err := strconv.Atoi(temp)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	tw, _ := taskwarrior.NewTaskWarrior("~/.taskrc")
	tw.FetchAllTasks()

	
	task, err := current_next.GetTaskWithId(tw, int32(task_id))
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	output := fmt.Sprintf("%d | %s", task_id, task.Description)
	fmt.Println(output)
}








