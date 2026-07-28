package main

import (
	"context"
	"fmt"
)

type ctxKey string

const userID ctxKey = "userID"

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, userID, "fadni")

	mm(ctx)
}

func mm(ctx context.Context) {
	getV(ctx)
}

func getV(ctx context.Context) {
	fmt.Printf("%s\n", ctx.Value(userID))
}