package main

import (
	"fmt"
	"os"
	
	"github.com/parth/go-learnings/week-06/github-repo-manager/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
