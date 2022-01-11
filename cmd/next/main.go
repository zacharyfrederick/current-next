package main

import (
	"strings"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"github.com/jubnzv/go-taskwarrior"
)

func get_task_with_id(t *taskwarrior.TaskWarrior, id int32) (*taskwarrior.Task, error) {
	for _, task := range t.Tasks {
		if task.Id == id {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("A task with that id could not be found")
}

func main() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Your home directory is not set")
		os.Exit(3)
	}
	
	data_path := filepath.Join(homedir, ".current-next", "next.data")
	if _, err := os.Stat(data_path); os.IsNotExist(err) {
		fmt.Println("The next task has not been set") 
		os.Exit(3)
	}

	data, err := os.ReadFile(data_path)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	temp := strings.Replace(string(data), "\n", "", -1)
	if temp == "" {
		fmt.Println("No next task set")
		os.Exit(0)
	}
	task_id, err := strconv.Atoi(temp)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	tw, _ := taskwarrior.NewTaskWarrior("~/.taskrc")
	tw.FetchAllTasks()

	
	task, err := get_task_with_id(tw, int32(task_id))
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	output := fmt.Sprintf("%d | %s", task_id, task.Description)
	fmt.Println(output)
}








