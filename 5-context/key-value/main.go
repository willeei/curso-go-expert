package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "123456")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println("Token:", token)
}
