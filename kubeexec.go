package main

import (
	"kubeexec/term"
)

func main() {
	err := term.RunScenario()
	if err != nil {
		panic(err)
	}
}
