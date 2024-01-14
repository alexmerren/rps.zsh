package main

import (
	"context"
	"os"

	"github.com/alexmerren/rps/src"
)

func main() {
	ctx := context.Background()
	token := os.Args[1]
	client := src.NewClient(ctx, token)
	client.ListUserRepos()
}
