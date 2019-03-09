package main

import (
    "fmt"
    "github.com/octokit/go-octokit/octokit"
)

func main() {
    client := octokit.NewClient(nil)

    url, err := octokit.UserURL.Expand(octokit.M{"user": "derrickwilliams"})
    if err != nil  {
        panic(fmt.Sprintf("URL error: %s", err.Error()))
    }

    user, result := client.Users(url).One()
    if result.HasError() {
        panic(fmt.Sprintf("URL error: %s", result.Error()))
    }

    fmt.Println(user.ReposURL)
}
