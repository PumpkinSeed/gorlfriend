package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	var behavior = getBehavior()
	if behavior {
		var (
			cmdOut []byte
			err    error
		)
		if cmdOut, err = exec.Command("go", os.Args[1:]...).Output(); err != nil {
			fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
			os.Exit(1)
		}
		fmt.Println(string(cmdOut))
		return
	}
	fmt.Println(reasons[randInt(0, len(reasons)-1)])
	return
}

func getBehavior() bool {
	if randInt(1, 20) > 5 {
		return false
	}
	return true
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

var reasons = []string{
	"I would rather go to cinema. :/",
	"I don't really want to do that.",
	"I'm on my period, sorry darling.",
	"Nothing",
	"Because you can't even solve THIS.",
	"Why are you mad?",
	"You’re hilarious!",
	"You're not good enough!",
	"You don't deserve me.",
	"You don't understand me!",
	"NO!",
}
