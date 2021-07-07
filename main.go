package main

import (
	"github.com/gleich/lumber"
	"github.com/gleich/solar/pkg/cmd"
)

func main() {
	err := cmd.RootCMD.Execute()
	if err != nil {
		lumber.Fatal(err, "Failed to execute root command")
	}
}
