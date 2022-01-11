package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Missing task id")
		os.Exit(3)
	}
	task_id := os.Args[1]

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

	f, err := os.Create(data_path)
	if err != nil {
		fmt.Println("Could not open the data file")
		os.Exit(3)
	}

	_, err = f.WriteString(string(task_id))
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	fmt.Printf("Next task set to %s\n", task_id)
}

