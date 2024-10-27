package main

import (
	"backend/app/awa/fetchers"
	"context"
	"fmt"
	"github.com/carlmjohnson/requests"
)

func main() {
	githubId := "fuhaoda"
	data := make([]fetchers.ReposFull, 0, 1)
	err := requests.URL("https://api.github.com/users/" + githubId + "/repos").
		ToJSON(&data).
		Fetch(context.Background())
	fmt.Printf("%v\n", err)
	for _, v := range data {
		if v.Fork == true {
			fmt.Printf("%v\n", v.FullName)
			fmt.Printf("%v\n", v.Parent.FullName)
			if v.Parent.FullName == "" {
				fmt.Printf("error!!!\n")
			}
		}
	}

}
