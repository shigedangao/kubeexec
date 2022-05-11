package main

import (
	"fmt"
	"kubeexec/term"
)

func main() {
	err := term.RunScenario()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
