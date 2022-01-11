package main

import (
	"os"
	"fmt"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Missing task id")
		os.Exit(3)
	}
	
	out, _ := exec.Command("task", os.Args[1], "done").CombinedOutput()
	fmt.Printf("%s", out)

}
