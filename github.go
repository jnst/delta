package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func sample() {
	client := github.NewClient(nil)
	ctx := context.Background()
	commits, _, err := client.Repositories.ListCommits(ctx, "etherdelta", "etherdelta.github.io", nil)
	if err != nil {
		println(err)
		return
	}
	for i := 0; i < len(commits); i++ {
		fmt.Printf("%v: %v\n", *commits[i].SHA, *commits[i].Commit.Message)
	}

	//notifications, _, err := client.Activity.ListRepositoryNotifications(ctx, "etherdelta", "etherdelta.github.io", nil)
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//
	//for i := 0; i < len(notifications); i++ {
	//	n := *notifications[i]
	//	fmt.Printf("id=%v, repo=%v, subject=%v\n", n.ID, n.Repository.Name, n.Subject.Title)
	//}

}
