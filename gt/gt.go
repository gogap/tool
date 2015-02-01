package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("must have commit message")
		return
	}
	var message string
	for k, v := range os.Args {
		if k > 0 {
			message += v + " "
		}
	}
	b, err := exec.Command("git", "pull").Output()
	if err != nil {
		fmt.Printf("git status error %s\n", err.Error())
		return
	}
	fmt.Println(string(b))

	b, err = exec.Command("git", "status").Output()
	if err != nil {
		fmt.Printf("git status error %s\n", err.Error())
		return
	}
	fmt.Println(string(b))

	err = exec.Command("git", "add", "--all").Run()
	if err != nil {
		fmt.Printf("git add error %s\n", err.Error())
		return
	}

	err = exec.Command("git", "commit", "-m", message).Run()
	if err != nil {
		fmt.Printf("git commit error %s\n", err.Error())
		return
	}

	err = exec.Command("git", "push").Run()
	if err != nil {
		fmt.Printf("git push error %s\n", err.Error())
		return
	}
}
