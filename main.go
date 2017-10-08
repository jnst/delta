package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func main() {
	client := github.NewClient(nil)
	ctx := context.Background()
	commits, _, err := client.Repositories.ListCommits(ctx, "etherdelta", "etherdelta.github.io", nil)
	if err != nil {
		println(err)
	}
	for i := 0; i < len(commits); i++ {
		fmt.Printf("%v: %v\n", *commits[i].SHA, *commits[i].Commit.Message)
	}
}
